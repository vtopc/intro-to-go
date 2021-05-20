package main

import (
	"testing"
)

func BenchmarkWithoutTrace(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//
	}
}
