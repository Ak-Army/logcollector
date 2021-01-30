package loki

import (
	"bytes"
	"net/http"
	"time"

	"github.com/Ak-Army/httpClient"
	"github.com/Ak-Army/httpClient/decoder"
	"github.com/Ak-Army/httpClient/middleware"
	"github.com/Ak-Army/xlog"

	"github.com/Ak-Army/logcollector/proto/loki"
)

type Client struct {
	client *httpClient.Client
	log    xlog.Logger
}

func NewClient(log xlog.Logger) *Client {
	c := &Client{
		client: httpClient.New(),
		log:    log,
	}
	c.client.Base("http://localhost:3100").
		Middleware(middleware.NewLoggerWrapper(log)).
		Middleware(middleware.NewTimeoutWrapper(120 * time.Second)).
		Middleware(middleware.NewRetryWrapper(3, func(request *http.Request, response *http.Response, err error) bool {
			if response != nil && response.StatusCode == 429 {
				log.Info("Retry")
				time.Sleep(2 * time.Second)
				return true
			}
			return false
		})).
		Middleware(middleware.NewResponseCodeWrapper(200, 299)).
		Middleware(NewRateLimitWrapper(time.Second, 2000000))
	return c
}

func (c *Client) Send(e *Entry) (*http.Response, error) {
	be := make(batchEntries)
	fp := e.Labels.String()
	be[fp] = &loki.Stream{
		Labels:  fp,
		Entries: []loki.Entry{e.Entry},
	}
	return c.send(be)
}

func (c *Client) send(b batchEntries) (*http.Response, error) {
	errors := &bytes.Buffer{}
	defer func() {
		xlog.Debug("response: ", errors.String())
	}()
	cc := c.client.Clone().ResponseDecoder(&decoder.Plain{}).BodyProvider(b).Post("/loki/api/v1/push")
	req, err := cc.Request()
	if err != nil {
		return nil, err
	}

	return cc.Do(req, errors, errors)
}
