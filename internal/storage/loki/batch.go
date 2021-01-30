package loki

import (
	"bytes"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"

	"github.com/Ak-Army/logcollector/proto/loki"
)

type batchEntries map[string]*loki.Stream

func (b batchEntries) ContentType() string {
	return "application/x-protobuf"
}

func (b batchEntries) Body() (io.Reader, error) {
	req := loki.PushRequest{
		Streams: make([]*loki.Stream, 0, len(b)),
	}
	for _, stream := range b {
		req.Streams = append(req.Streams, stream)
	}
	buf, err := proto.Marshal(&req)
	if err != nil {
		return nil, err
	}
	buf = snappy.Encode(nil, buf)
	return bytes.NewBuffer(buf), nil
}

type batchEntriesSorter func(a, b loki.Entry) bool

type batchEntriesSortable struct {
	values     []loki.Entry
	size       int
	comparator batchEntriesSorter
}

func (s batchEntriesSortable) Len() int {
	return s.size
}

func (s batchEntriesSortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s batchEntriesSortable) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j])
}

var timeSort = func(e1, e2 loki.Entry) bool {
	if e1.Timestamp.Before(e2.Timestamp) {
		return true
	}
	return false
}
