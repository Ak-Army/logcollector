package influxdb

import (
	"fmt"
	"strings"
	"time"

	"github.com/Ak-Army/logcollector/internal/storage"

	"github.com/Ak-Army/xlog"
	client "github.com/influxdata/influxdb1-client/v2"
	"go.uber.org/atomic"
)

type batchClient struct {
	client         client.Client
	batch          client.BatchPoints
	batchWait      time.Duration
	batchTimer     *time.Timer
	batchSize      int
	maxSize        int
	log            xlog.Logger
	entriesChannel chan pointWithSize
	done           chan struct{}
	sent           atomic.Int64
}

type pointWithSize struct {
	*client.Point
	size int
}

func New(log xlog.Logger, entryBufferSize int, batchSize int, batchWait time.Duration) storage.Storage {
	var err error
	c := &batchClient{
		log:            log,
		maxSize:        batchSize,
		batchWait:      batchWait,
		entriesChannel: make(chan pointWithSize, entryBufferSize),
		done:           make(chan struct{}),
	}
	if c.client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "admin",
		Password: "admin",
	}); err != nil {
		log.Error("Unable to connect", err)
		return nil
	}
	if err := c.newBatchPoints(); err != nil {
		log.Error("Unable to create batch points", err)
		return nil
	}
	go c.run()

	return c
}

func (c *batchClient) DropDatabase() error {
	query := client.NewQuery(fmt.Sprintf(`DROP DATABASE "%s"`, "log"), "", "")
	c.log.Debugf("Drop database: %s", "log")
	_, err := c.client.Query(query)
	return err
}

func (c *batchClient) DropApp(app string) error {
	query := client.NewQuery(fmt.Sprintf(`DROP MEASUREMENT "%s"`, app), "log", "")
	c.log.Debugf("DROP MEASUREMENT %s", app)
	_, err := c.client.Query(query)
	return err
}

func (c *batchClient) DeleteByDate(app string, dateFrom, dateTo time.Time) error {
	query := client.NewQuery(
		fmt.Sprintf(`DELETE FROM "%s" WHERE time < '%s' and time > '%s'`,
			app,
			dateFrom.Format("20060102"),
			dateTo.Format("20060102"),
		),
		"log",
		"",
	)
	c.log.Debugf("Delete by date database: %s %s->%s", app, dateFrom, dateTo)
	if _, err := c.client.Query(query); err != nil {
		c.log.Error("Unable to create database", err)
		return err
	}
	return nil
}

func (c *batchClient) Send(line storage.LogLine) error {
	p, err := client.NewPoint(line.App, line.Tags, line.Fields, line.Time)
	if err != nil {
		return err
	}
	c.entriesChannel <- pointWithSize{p, line.Size}
	return nil
}

func (c *batchClient) Stop() error {
	close(c.entriesChannel)
	<-c.done
	return c.client.Close()
}

func (c *batchClient) AlreadySent() int64 {
	return c.sent.Swap(0)
}

func (c *batchClient) run() {
	c.batchTimer = time.NewTimer(c.batchWait)
	defer func() {
		if c.batch != nil && len(c.batch.Points()) > 0 {
			if c.write() {
				c.log.Warn("unable to send data")
			}
		}
		close(c.done)
	}()
	currSize := 0
	for {
		select {
		case ll, ok := <-c.entriesChannel:
			if !ok {
				return
			}
			currSize += ll.size
			if currSize > c.maxSize {
				if c.write() {
					c.log.Warn("unable to send data")
				}
				currSize = ll.size
			}
			c.batch.AddPoint(ll.Point)
		case <-c.batchTimer.C:
			if len(c.batch.Points()) > 0 {
				if c.write() {
					c.log.Warn("unable to send data")
				}
				currSize = 0
			}
			c.batchTimer.Reset(c.batchWait)
		}
	}
}

func (c *batchClient) newBatchPoints() error {
	var err error
	if c.batch, err = client.NewBatchPoints(client.BatchPointsConfig{
		Database: "log",
	}); err != nil {
		return err
	}
	return nil
}

func (c *batchClient) write() bool {
	c.sent.Add(int64(len(c.batch.Points())))
	err := c.client.Write(c.batch)
	if err != nil {
		if strings.Contains(err.Error(), "database not found") {
			query := client.NewQuery(fmt.Sprintf(`CREATE DATABASE "%s"`, "log"), "", "")
			c.log.Debugf("Create database: %s", "log")
			if _, err := c.client.Query(query); err != nil {
				c.log.Error("Unable to create database", err)
				return true
			}
		} else {
			c.log.Error("Unable to push write data", err)
		}
	}
	if err := c.newBatchPoints(); err != nil {
		c.log.Error("Unable to create batch points", err)
		return true
	}
	c.batchTimer.Reset(c.batchWait)
	return false
}
