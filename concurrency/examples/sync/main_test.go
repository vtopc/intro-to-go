package main

import (
	"io"
	"log"
	"testing"

	"concurrency"
)

var requests = concurrency.GenerateRequests(concurrency.Count)

func init() {
	log.SetOutput(io.Discard)
}

func BenchmarkSync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoSync(requests)
	}
}
