package loki

import (
	"sort"
	"time"

	"github.com/Ak-Army/xlog"
	"github.com/uber-go/atomic"

	"github.com/Ak-Army/logcollector/internal/storage"
	"github.com/Ak-Army/logcollector/proto/loki"
)

type batchClient struct {
	client         *Client
	batch          batchEntries
	batchWait      time.Duration
	batchTimer     *time.Timer
	batchSize      int
	maxSize        int
	log            xlog.Logger
	entriesChannel chan *Entry
	done           chan interface{}
	sent           atomic.Int64
}

func New(log xlog.Logger, entryBufferSize int, batchSize int, batchWait time.Duration) storage.Storage {
	bc := &batchClient{
		client:         NewClient(log),
		log:            log,
		maxSize:        batchSize,
		batchWait:      batchWait,
		entriesChannel: make(chan *Entry, entryBufferSize),
		done:           make(chan interface{}),
		batch:          make(batchEntries),
	}
	go bc.run()

	return bc
}

func (c *batchClient) Send(line storage.LogLine) error {
	e := &Entry{
		Labels: Labels{},
		Entry: loki.Entry{
			Timestamp: line.Time,
			Line:      line.Fields["raw"].(string),
		},
	}
	e.Labels.Add("app", line.App)
	var keys []string
	for k, _ := range line.Tags {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		e.Labels.Add(k, line.Tags[k])
	}
	c.entriesChannel <- e
	return nil
}

func (c *batchClient) DropDatabase() error {
	return nil
}

func (c *batchClient) DropApp(app string) error {
	return nil
}

func (c *batchClient) DeleteByDate(app string, dateFrom, dateTo time.Time) error {
	return nil
}

func (c *batchClient) Stop() error {
	close(c.entriesChannel)
	<-c.done
	return nil
}

func (c *batchClient) run() {
	c.batchTimer = time.NewTimer(c.batchWait)
	defer func() {
		if c.batch != nil && len(c.batch) > 0 {
			c.write()
		}
		close(c.done)
	}()

	for {
		select {
		case ll, ok := <-c.entriesChannel:
			if !ok {
				return
			}
			fp := ll.Labels.String()
			stream, ok := c.batch[fp]
			if !ok {
				stream = &loki.Stream{
					Labels: fp,
				}
				c.batch[fp] = stream
			}
			stream.Entries = append(stream.Entries, ll.Entry)
			c.batchSize += len(ll.Line)
			if c.batchSize > c.maxSize {
				c.write()
			}
		case <-c.batchTimer.C:
			if len(c.batch) > 0 {
				c.write()
			}
			c.batchTimer.Reset(c.batchWait)
		}
	}
}

func (c *batchClient) write() {
	for _, stream := range c.batch {
		sort.Sort(batchEntriesSortable{values: stream.Entries, size: len(stream.Entries), comparator: timeSort})
	}
	if _, err := c.client.send(c.batch); err != nil {
		c.log.Error("Batch send error: ", err)
	}
	c.batchSize = 0
	c.batch = make(batchEntries)
	c.batchTimer.Reset(c.batchWait)
}
