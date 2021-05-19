package main

import (
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
)

var tracedClient = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &ochttp.Transport{
		// Use Google Cloud propagation format.
		Propagation:    &propagation.HTTPFormat{},
		FormatSpanName: formatSpanName,
	},
}
