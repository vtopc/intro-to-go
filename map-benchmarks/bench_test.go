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

func BenchmarkIfaceMapSearch(b *testing.B) {
	m := newIfaceMap()

	for n := 0; n < b.N; n++ {
		_ = m["E"]
	}
}

func BenchmarkIntMapSearch(b *testing.B) {
	m := newIntMap()

	for n := 0; n < b.N; n++ {
		_ = m["E"]
	}
}

func BenchmarkSwitchSearch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = switchSearch("E")
	}
}

func BenchmarkStructSearch(b *testing.B) {
	s := newStruct()

	for n := 0; n < b.N; n++ {
		_ = s.E
	}
}
