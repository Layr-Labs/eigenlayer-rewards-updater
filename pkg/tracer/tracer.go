package tracer

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func StartTracer(enabled bool) {
	if !enabled {
		mocktracer.Start()
		return
	}
	ddTracer.Start(
		ddTracer.WithGlobalServiceName(true),
		ddTracer.WithDebugMode(false),
		ddTracer.WithLogStartup(false),
	)
}
