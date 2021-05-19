package main

import (
	"log"
	"net/http"
	"strings"

	"go.opencensus.io/trace"
)

const healthzPath = "/healthz"

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "test-tracing/handler")
	defer span.End()
	w.Header().Set("X-Trace-Id", span.SpanContext().TraceID.String())

	// The trace ID from the incoming request will be
	// propagated to the outgoing request.
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://www.google.com",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// The outgoing request will be traced with r's trace ID.
	resp, err := tracedClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Because we don't read the resp.Body, need to manually call Close().
	_ = resp.Body.Close()

	_, _ = w.Write([]byte("bar"))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("OK"))
}

func isHealthEndpoint(r *http.Request) bool {
	if strings.HasSuffix(r.URL.String(), healthzPath) {
		return true
	}

	return false
}
