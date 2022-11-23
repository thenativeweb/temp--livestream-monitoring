package metrics

import (
	"sync"

	"github.com/thenativeweb/monitoring/maps"
)

type HttpRequests struct {
	mutex    sync.RWMutex
	counters map[string]int64
}

func NewHttpRequests() *HttpRequests {
	return &HttpRequests{
		counters: map[string]int64{},
	}
}

func (httpRequests *HttpRequests) RegisterRoute(route string) {
	httpRequests.mutex.Lock()
	defer httpRequests.mutex.Unlock()

	if _, ok := httpRequests.counters[route]; ok {
		panic("route already registered")
	}

	httpRequests.counters[route] = 0
}

func (httpRequests *HttpRequests) Increment(route string) {
	httpRequests.mutex.Lock()
	defer httpRequests.mutex.Unlock()

	if _, ok := httpRequests.counters[route]; !ok {
		panic("route not registered")
	}

	httpRequests.counters[route]++
}

func (httpRequests *HttpRequests) Value() map[string]int64 {
	httpRequests.mutex.RLock()
	defer httpRequests.mutex.RUnlock()

	counters := maps.Clone(httpRequests.counters)

	return counters
}
