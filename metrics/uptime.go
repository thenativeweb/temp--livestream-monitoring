package metrics

import (
	"time"
)

type Uptime struct {
	startTime time.Time
}

func NewUptime(startTime time.Time) *Uptime {
	return &Uptime{
		startTime: startTime,
	}
}

func (uptime *Uptime) Value() (int64, string) {
	return time.Since(uptime.startTime).Milliseconds(), Milliseconds
}
