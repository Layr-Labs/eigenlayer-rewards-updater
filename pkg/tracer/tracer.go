package tracer

import (
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"os"
)

func StartTracer() {
	os.Setenv("DD_TRACE_STARTUP_LOGS", "false")
	ddTracer.Start(
		ddTracer.WithGlobalServiceName(true),
		ddTracer.WithDebugMode(false),
	)
}
