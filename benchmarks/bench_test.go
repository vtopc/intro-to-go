package main

import (
	_ "embed"
	"testing"
)

const count = 100

var s []S

func BenchmarkIndexSlice(b *testing.B) {
	var s []S

	for n := 0; n < b.N; n++ {
		s = make([]S, count)

		for i := 0; i < count; i++ {
			s[i] = newStruct()
		}
	}

	_ = s
}

func BenchmarkAppendSlice(b *testing.B) {
	var s []S

	for n := 0; n < b.N; n++ {
		s = make([]S, 0, count)

		for i := 0; i < count; i++ {
			s = append(s, newStruct())
		}
	}

	_ = s
}
