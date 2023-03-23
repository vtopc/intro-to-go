package main

import (
	"context"
	"io"
	"log"
	"testing"

	"concurrency"
)

var requests = concurrency.GenerateRequests(concurrency.Count)

func init() {
	log.SetOutput(io.Discard)
}

func BenchmarkConcPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoAsync(context.TODO(), requests)
	}
}
