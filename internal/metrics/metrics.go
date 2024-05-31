package metrics

import (
	"github.com/DataDog/datadog-go/v5/statsd"
)

func NewStatsdClient(addr string) (*statsd.Client, error) {
	// if the addr is empty, statsd will look at the envvar DD_DOGSTATSD_URL
	return statsd.New(addr)
}
