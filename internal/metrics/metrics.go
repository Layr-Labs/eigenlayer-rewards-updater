package metrics

import (
	"errors"
	"github.com/DataDog/datadog-go/v5/statsd"
)

var statsdClient statsd.ClientInterface

func InitStatsdClient(addr string, enabled bool) (statsd.ClientInterface, error) {
	if !enabled {
		statsdClient = &statsd.NoOpClient{}
		return statsdClient, nil
	}
	var err error
	statsdClient, err = statsd.New(addr,
		statsd.WithNamespace("eigenlayer_rewards_updater."),
		statsd.WithMaxMessagesPerPayload(1),
	)
	return statsdClient, err
}

func GetStatsdClient() statsd.ClientInterface {
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
