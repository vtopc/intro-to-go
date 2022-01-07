package main

import (
	"testing"
)

func BenchmarkIfaceMapAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newIfaceMap()
	}
}

func BenchmarkIntMapAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newIntMap()
	}
}

func BenchmarkStructAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newStruct()
	}
}

func BenchmarkIntMapSearch(b *testing.B) {
	m := newIntMap()

	for n := 0; n < b.N; n++ {
		_ = m["E"]
	}
}

func BenchmarkStructSearch(b *testing.B) {
	s := newStruct()

	for n := 0; n < b.N; n++ {
		_ = s.E
	}
}
