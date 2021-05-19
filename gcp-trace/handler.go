package main

import (
	"log"
	"net/http"

	"go.opencensus.io/trace"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "handler")
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
