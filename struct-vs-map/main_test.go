package main

import (
	"testing"
)

type S struct {
	A int
	B int
	C int
	D int
	E int
	F int
	G int
	H int
	J int
	K int
}

func BenchmarkMapAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newMap()
	}
}

func BenchmarkStructAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newStruct()
	}
}

func BenchmarkMapSearch(b *testing.B) {
	m := newMap()

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

func newMap() map[string]interface{} {
	return map[string]interface{}{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"J": 9,
		"K": 10,
	}
}

func newStruct() S {
	return S{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
		F: 6,
		G: 7,
		H: 8,
		J: 9,
		K: 10,
	}
}
