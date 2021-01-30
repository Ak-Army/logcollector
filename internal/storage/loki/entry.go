package loki

import (
	"strings"
	"unsafe"

	"github.com/Ak-Army/logcollector/proto/loki"
)

type Entry struct {
	Labels
	loki.Entry
}

type Labels struct {
	buf []byte
}

func (l *Labels) Add(key, value string) *Labels {
	if strings.Contains(value, "{") {
		return l
	}
	l.buf = append(l.buf, ',')
	l.buf = append(l.buf, key...)
	l.buf = append(l.buf, '=')
	l.buf = append(l.buf, '"')
	l.buf = append(l.buf, value...)
	l.buf = append(l.buf, '"')

	return l
}

func (l *Labels) AddByte(key, value []byte) *Labels {
	if strings.Contains(string(value), "{") {
		return l
	}
	l.buf = append(l.buf, ',')
	l.buf = append(l.buf, key...)
	l.buf = append(l.buf, '=')
	//l.buf = append(l.buf, '"')
	l.buf = append(l.buf, value...)
	//l.buf = append(l.buf, '"')

	return l
}
func (l Labels) String() string {
	if len(l.buf) == 0 {
		return ""
	}
	return "{" + (*(*string)(unsafe.Pointer(&l.buf)))[1:] + "}"
}
