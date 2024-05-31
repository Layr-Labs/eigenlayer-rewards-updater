package metrics

import (
	"errors"
	"github.com/DataDog/datadog-go/v5/statsd"
)

var statsdClient *statsd.Client

func InitStatsdClient(addr string) (*statsd.Client, error) {
	var err error
	statsdClient, err = statsd.New(addr)
	return statsdClient, err
}

func GetStatsdClient() *statsd.Client {
	if statsdClient == nil {
		panic(errors.New("statsd client not initialized"))
	}
	return statsdClient
}

const (
	Counter_UpdateRuns     = "update_runs"
	Counter_UpdateFails    = "update_fails"
	Counter_UpdateSuccess  = "update_success"
	Counter_UpdateNoUpdate = "update_no_update"
)
