// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

const tracePrefix = "test-tracing"

func main() {
	registerTrace()

	router := NewRouter()
	traceRouter := TraceWrapper(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Listening on %s", addr)

	server := NewServer(addr, traceRouter)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func NewRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/foo", http.HandlerFunc(handler))
	router.Handle("/healthz", http.HandlerFunc(healthz))

	return router
}

func TraceWrapper(h http.Handler) http.Handler {
	// Use an ochttp.Handler in order to instrument OpenCensus for incoming
	// requests.
	w := &ochttp.Handler{
		// Use the Google Cloud propagation format.
		Propagation:      &propagation.HTTPFormat{},
		Handler:          h,
		IsHealthEndpoint: isHealthEndpoint,
		FormatSpanName:   formatSpanName,
	}

	return w
}

func formatSpanName(r *http.Request) string {
	return tracePrefix + r.URL.Path
}

func registerTrace() {
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