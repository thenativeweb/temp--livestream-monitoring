package metrics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/monitoring/metrics"
)

func TestUnits(t *testing.T) {
	assert.Equal(t, "milliseconds", metrics.Milliseconds)
}
