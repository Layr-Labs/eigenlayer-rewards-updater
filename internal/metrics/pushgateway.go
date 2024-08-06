package metrics

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	prometheusPusherClient *PrometheusPusherClient
)

const (
	CounterUpdateRunsInvoked  = "invoked"
	CounterUpdateRunsFailed   = "failed"
	CounterUpdateRunsSuccess  = "success"
	CounterUpdateRunsNoUpdate = "no-update"
)

// PrometheusPusherClient is a pusher for the metrics to the pushgateway.
type PrometheusPusherClient struct {
	pushgatewayAddr string
	pusher          *push.Pusher
	// Metrics
	counterUpdateRuns *prometheus.CounterVec
}

// InitPrometheusPusherClient creates a new PrometheusPusherClient. The client will
// try to fetch the current metrics values from the pushgateway and initialize it
// with them. This is useful to maintain a good history of the metrics. If the metric
// is not found, it will be initialized with 0.
func InitPrometheusPusherClient(addr string, jobName string) error {
	prometheusPusherClient = &PrometheusPusherClient{
		pushgatewayAddr: addr,
		pusher:          push.New(addr, jobName),
		counterUpdateRuns: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "counter_update_run_total",
		}, []string{"status"}),
	}
	prometheusPusherClient.pusher.Collector(prometheusPusherClient.counterUpdateRuns)
	return prometheusPusherClient.initCounterMetrics()
}

// IncCounterUpdateRun increments the counter_update_runs metric with the given status.
// If the Pushgateway client is not initialized, the call is ignored. To avoid ignoring
// make sure to call the InitPrometheusPusherClient function before using this function.
// The status can be one of the following: "invoked", "failed", "success", "no-update"
func IncCounterUpdateRun(status string) {
	if prometheusPusherClient == nil {
		return
	}
	prometheusPusherClient.incCounterUpdateRun(status)
}

// PushToPushgateway pushes the metrics to the pushgateway.
// If the Pushgateway client is not initialized, the call is ignored. To avoid ignoring
// make sure to call the InitPrometheusPusherClient function before using this function.
func PushToPushgateway() error {
	if prometheusPusherClient == nil {
		return nil
	}
	return prometheusPusherClient.push()
}

func (p *PrometheusPusherClient) incCounterUpdateRun(status string) {
	p.counterUpdateRuns.WithLabelValues(status).Inc()
}

func (p *PrometheusPusherClient) push() error {
	return p.pusher.Push()
}

func (p *PrometheusPusherClient) initCounterMetrics() error {
	resp, err := http.Get(p.pushgatewayAddr + "/metrics")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch metrics from pushgateway: %s", resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyString := string(body)
	initializers := []struct {
		regex   *regexp.Regexp
		counter prometheus.Counter
	}{
		{
			regex:   regexp.MustCompile(`counter_update_run_total{.*\bstatus="failed"} (\d+)`),
			counter: p.counterUpdateRuns.WithLabelValues(CounterUpdateRunsFailed),
		},
		{
			regex:   regexp.MustCompile(`counter_update_run_total{.*\bstatus="success"} (\d+)`),
			counter: p.counterUpdateRuns.WithLabelValues(CounterUpdateRunsSuccess),
		},
		{
			regex:   regexp.MustCompile(`counter_update_run_total{.*\bstatus="invoked"} (\d+)`),
			counter: p.counterUpdateRuns.WithLabelValues(CounterUpdateRunsInvoked),
		},
		{
			regex:   regexp.MustCompile(`counter_update_run_total{.*\bstatus="no-update"} (\d+)`),
			counter: p.counterUpdateRuns.WithLabelValues(CounterUpdateRunsNoUpdate),
		},
	}
	for _, initializer := range initializers {
		matches := initializer.regex.FindStringSubmatch(bodyString)
		if len(matches) > 0 {
			value, err := strconv.ParseFloat(matches[1], 64)
			if err != nil {
				return err
			}
			initializer.counter.Add(value)
		}
	}

	return nil
}
