package loki

import (
	"net/http"
	"time"

	"github.com/Ak-Army/httpClient/middleware"
)

type rateLimitDoer struct {
	doer         middleware.Doer
	timeout      time.Duration
	rateLimit    int64
	now          time.Time
	requestsSize int64
}

func NewRateLimitWrapper(timeout time.Duration, rateLimit int64) middleware.Wrapper {
	td := rateLimitDoer{
		timeout:      timeout,
		rateLimit:    rateLimit,
		requestsSize: 0,
		now:          time.Now(),
	}
	return func(doer middleware.Doer) middleware.Doer {
		td.doer = doer
		return td.Do
	}
}

func (d *rateLimitDoer) Do(req *http.Request, successV, failureV interface{}) (*http.Response, error) {
	now := time.Now()
	if d.now.Add(d.timeout).Before(now) {
		d.requestsSize = 0
	}
	d.requestsSize += req.ContentLength
	if d.rateLimit <= d.requestsSize {
		time.Sleep(d.timeout - now.Sub(d.now))
	}
	d.now = now
	return d.doer(req, successV, failureV)
}
