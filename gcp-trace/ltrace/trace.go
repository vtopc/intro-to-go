package ltrace

import (
	"log"
	"net/http"
	"os"

	"gcp-trace/httpserver/healthz"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

const Prefix = "test-tracing"

func FormatSpanName(r *http.Request) string {
	return Prefix + r.URL.Path
}

func RegisterTrace() {
	// Create and register a OpenCensus Stackdriver Trace exporter.
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)

	// By default, traces will be sampled relatively rarely. To change the
	// sampling frequency for your entire program, call ApplyConfig. Use a
	// ProbabilitySampler to sample a subset of traces, or use AlwaysSample to
	// collect a trace on every run.
	//
	// Be careful about using trace.AlwaysSample in a production application
	// with significant traffic: a new trace will be started and exported for
	// every request.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func HTTPHandlerWrapper(h http.Handler) http.Handler {
	// Use an ochttp.Handler in order to instrument OpenCensus for incoming
	// requests.
	return &ochttp.Handler{
		// Use the Google Cloud propagation format.
		Propagation:      &propagation.HTTPFormat{},
		Handler:          h,
		IsHealthEndpoint: healthz.IsHealthEndpoint,
		FormatSpanName:   FormatSpanName,
	}
}
