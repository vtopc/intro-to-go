package main

import (
	"io/ioutil"
	"log"
	"testing"
)

var requests = generateRequest(Count)

func init() {
	log.SetOutput(ioutil.Discard)
}

func BenchmarkDoAsync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoAsync(requests)
	}
}

func BenchmarkDoSync(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoSync(requests)
	}
}
