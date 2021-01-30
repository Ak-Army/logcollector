package cmd

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Ak-Army/logcollector/internal/ssh_client"
	"github.com/Ak-Army/logcollector/internal/storage"
	"github.com/Ak-Army/logcollector/internal/storage/influxdb"
	"github.com/Ak-Army/logcollector/internal/storage/loki"

	"github.com/Ak-Army/cli"
	"github.com/Ak-Army/xlog"
	"github.com/go-logfmt/logfmt"
	"github.com/sgreben/flagvar"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var apps = []string{}

func init() {
	cli.RootCommand().AddCommand("collect", &Collect{})
}

type Collect struct {
	Apps            flagvar.Strings `flag:"apps, app name"`
	Servers         flagvar.Strings `flag:"servers, app name"`
	FromServer      string          `flag:"fs, from server"`
	FromApp         string          `flag:"fa, from app"`
	Date            string          `flag:"date, date"`
	DropDb          bool            `flag:"dropDB, drop db"`
	DropMeasurement bool            `flag:"dropMeas, drop measurement"`
	Loki            bool            `flag:"loki, send data to loki"`
	ctx             context.Context
	syslog          ssh_client.SSHClient
	fileProcess     chan string
}

func (c Collect) Help() string {
	return `Usage: log-collector <command> [command options] app_name`
}

func (c Collect) Synopsis() string {
	return "Collect logs from path"
}

func (c Collect) Run(ctx context.Context) error {
	fromApp := false
	if c.FromApp == "" {
		fromApp = true
	}
	if len(c.Apps.Values) == 0 {
		for _, app := range apps {
			if fromApp || c.FromApp == app {
				c.Apps.Set(app)
				fromApp = true
			}
		}
	}
	if len(c.Servers.Values) == 0 {
		c.Servers.Set("*")
	}
	if c.Date == "" {
		c.Date = time.Now().Add(time.Hour * -24).Format("20060102")
	}
	c.ctx = ctx
	sshConfig := &ssh.ClientConfig{
		User: "peter.hunyadvari",
		Auth: []ssh.AuthMethod{
			c.sshAgent(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	c.syslog = ssh_client.SSHClient{
		Config: sshConfig,
		Host:   "syslog-server",
		Port:   22,
	}
	maxDownloader := 1
	maxFileProcesor := 1
	download := make(chan string)
	c.fileProcess = make(chan string, 3)
	wg := sync.WaitGroup{}
	for i := 0; i < maxDownloader; i++ {
		go func() {
			for {
				path := <-download
				wg.Add(1)
				if !c.downloadFile(ctx, path) {
					download <- path
				}
				wg.Done()
			}
		}()
	}
	for i := 0; i < maxFileProcesor; i++ {
		go func() {
			for {
				path := <-c.fileProcess
				wg.Add(1)
				if !c.processFile(ctx, path) {
					c.fileProcess <- path
				}
				wg.Done()
			}
		}()
	}

	storage := c.storage()
	defer storage.Stop()
	if c.DropDb {
		if err := storage.DropDatabase(); err != nil {
			return err
		}
	}
	fromServer := false
	if c.FromServer == "" {
		fromServer = true
	}
	for _, app := range c.Apps.Values {
		var dirs []string
		for _, s := range c.Servers.Values {
			dirs = append(dirs, fmt.Sprintf(" /var/log/remote/%s/%s/%s/*", s, c.Date, app))
		}
		cmd := fmt.Sprintf("ls %s", strings.Join(dirs, " "))
		var cmdOut bytes.Buffer
		err := c.syslog.RunCommand(&ssh_client.SSHCommand{
			Path:   cmd,
			Env:    []string{},
			Stdin:  os.Stdin,
			Stdout: &cmdOut,
			Stderr: os.Stderr,
		})
		log := xlog.Copy(xlog.FromContext(ctx))
		if err == nil {
			if c.DropMeasurement {
				if err := storage.DropApp(app); err != nil {
					return err
				}
			}
			loc, _ := time.LoadLocation("Local")
			year, _ := strconv.Atoi(c.Date[0:3])
			month, _ := strconv.Atoi(c.Date[4:5])
			day, _ := strconv.Atoi(c.Date[6:7])
			dateFrom := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
			storage.DeleteByDate(app, dateFrom, dateFrom.Add(24*time.Hour))
			time.Sleep(10 * time.Millisecond)
			scanner := bufio.NewScanner(&cmdOut)
			for scanner.Scan() {
				path := scanner.Text()
				if fromServer || strings.HasPrefix(path, "/var/log/remote/"+c.FromServer) {
					fromServer = true
					download <- scanner.Text()
					continue
				}
				log.Infof("Skip: %s", path)
			}
			if err := scanner.Err(); err != nil {
				log.Error(err)
			}
		}
	}
	wg.Wait()
	return nil
}

var s storage.Storage

func (c Collect) storage() storage.Storage {
	if c.Loki {
		if s == nil {
			s = loki.New(xlog.FromContext(c.ctx), 1000, 2000000, 5*time.Second)
		}
		return s
	}
	return influxdb.New(xlog.FromContext(c.ctx), 1000, 10000000, 5*time.Second)
}

func (c Collect) sshAgent() ssh.AuthMethod {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		xlog.FromContext(c.ctx).Fatal(err)
	}
	return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
}

func (c Collect) downloadFile(ctx context.Context, path string) bool {
	log := xlog.Copy(xlog.FromContext(ctx))
	log.SetField("path", path)
	paths := strings.Split(path, "/")
	if len(paths) < 4 {
		log.Errorf("Unable to download file, %s", path)
		return false
	}
	file := fmt.Sprintf("%s_%s_%s.log", paths[4], paths[5], paths[7])
	if err := c.syslog.FileFromRemoteHost(file, path); err != nil {
		log.Error("Unable to download file")
		return false
	}
	c.fileProcess <- file

	return true
}

func (c Collect) processFile(ctx context.Context, file string) bool {
	log := xlog.Copy(xlog.FromContext(ctx))
	log.SetField("path", file)

	f, err := os.Open(file)
	if err != nil {
		log.Error("Unable to open file")
		return false
	}
	defer f.Close()
	db := c.storage()
	lineProcessor := 1
	line := make(chan string)
	wg := sync.WaitGroup{}
	for i := 0; i < lineProcessor; i++ {
		go func() {
			for {
				l := <-line
				if err := c.processLine(db, l, file); err != nil {
					log.Error(err)
					line <- l
					continue
				}
				wg.Done()
			}
		}()
	}
	scanner := bufio.NewScanner(f)
	sent := 0
	for scanner.Scan() {
		wg.Add(1)
		sent++
		line <- scanner.Text()
	}
	wg.Wait()
	if err := scanner.Err(); err != nil {
		log.Error(err)
	}
	//db.Stop()
	log.Infof("Sent: %d", sent)
	os.Remove(file)
	return true
}

func (c Collect) processLine(store storage.Storage, raw string, file string) error {
	logs := strings.SplitN(raw, " ", 4)
	if len(logs) < 4 {
		return errors.New("too short log line: " + raw)
	}

	ll := storage.LogLine{
		App:    strings.Split(logs[2], "[")[0],
		Tags:   make(map[string]string),
		Fields: make(map[string]interface{}),
		Time:   time.Now(),
		Size:   len(raw),
	}
	v := strings.Split(logs[0], "+")
	ll.Time, _ = time.Parse("2006-01-02T15:04:05.99999", v[0])
	ll.Tags["host"] = logs[1]
	ll.Fields["raw"] = logs[3]

	dec := logfmt.NewDecoder(strings.NewReader(logs[3]))
	for dec.ScanRecord() {
		for dec.ScanKeyval() {
			val := string(dec.Value())
			key := string(dec.Key())
			switch key {
			case "requestBody":
				continue
			case "customer":
				ll.Fields[key] = val
			case "method", "topic":
				ll.Tags["method_topic"] = val
			case "time":
				v := strings.Split(val, " m=")
				ll.Time, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", v[0])
			case "phone", "op", "secondaryProj", "queueId", "userName", "mode":
				ll.Fields[key] = val
			case "stat":
				if apps[0] == "go-queue" {
					//stat="QueueStat{running=0, completed=0, pending=0, total=0, lastJobTime=1970-01-01 01:00:00 +0100 CET, avWorker=1, busyWorker=0}"
					kvs := strings.Split(val[10:len(val)-2], ", ")
					for _, kv := range kvs {
						queueStat := strings.Split(kv, "=")
						if i, err := strconv.ParseInt(queueStat[1], 10, 64); err == nil {
							ll.Fields[queueStat[0]] = i
						}
					}
				}
				ll.Fields[key] = val
			default:
				if i, err := strconv.ParseInt(val, 10, 64); err == nil {
					ll.Fields[key] = i
				} else if i, err := strconv.ParseFloat(val, 64); err == nil {
					ll.Fields[key] = i
				} else if i, err := time.ParseDuration(val); err == nil {
					ll.Fields[key] = i.Milliseconds()
				} else {
					ll.Fields[key] = val
				}
			}
		}
	}
	return store.Send(ll)
}

// nginx logokhoz
/*quantile_over_time(0.99,
  {app="nginx_access"} |= "evcc_callback_proxy"
    | regexp `(?P<host>\S+) (?P<ip>\S+) (?P<ident>\S+) (?P<user>\S+) (?P<status>\S+) "(?P<url>[^"]+)" (?P<request_size>\S+) "(?P<aa>[^"]+)" "(?P<ss>[^"]+)" "(?P<dd>[^"]+)" "(?P<response_size>[^"]+)" "(?P<response_time>[^"]+)" "(?P<ff>[^"]+)" "(?P<gg>[^"]+)"`
    | unwrap response_time [1s]) by (app)
*/
