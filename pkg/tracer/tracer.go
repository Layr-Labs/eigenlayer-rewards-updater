package tracer

import (
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func StartTracer() {
	ddTracer.Start(
		ddTracer.WithGlobalServiceName(true),
		ddTracer.WithDebugMode(false),
	)
}
