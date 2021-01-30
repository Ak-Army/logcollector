package storage

import "time"

type LogLine struct {
	App    string
	Tags   map[string]string
	Fields map[string]interface{}
	Time   time.Time
	Size   int
}

type Storage interface {
	Send(LogLine) error
	Stop() error
	DropApp(app string) error
	DeleteByDate(app string, dateFrom time.Time, dateTo time.Time) error
	DropDatabase() error
}
