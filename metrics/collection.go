package metrics

import "time"

type Collection struct {
	Uptime       *Uptime
	HttpRequests *HttpRequests
}

func NewCollection(startTime time.Time) *Collection {
	return &Collection{
		Uptime:       NewUptime(startTime),
		HttpRequests: NewHttpRequests(),
	}
}
