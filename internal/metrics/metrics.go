package metrics

import (
	"github.com/DataDog/datadog-go/v5/statsd"
)

var statsdClient *statsd.Client

func GetStatsdClient(addr string) (*statsd.Client, error) {
	var err error
	if statsdClient != nil {
		return statsdClient, nil
	}
	// if the addr is empty, statsd will look at the envvar DD_DOGSTATSD_URL
	statsdClient, err = statsd.New(addr)
	return statsdClient, err
}
