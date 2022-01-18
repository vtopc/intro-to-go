package main

import (
	_ "embed"
	"testing"
)

const count = 100

func BenchmarkSliceIndex(b *testing.B) {
	var s []S

	for n := 0; n < b.N; n++ {
		s = make([]S, count)

		for i := 0; i < count; i++ {
			s[i] = newStruct()
		}
	}

	_ = s
}

func BenchmarkSliceAppend(b *testing.B) {
	var s []S

	for n := 0; n < b.N; n++ {
		s = make([]S, 0, count)

		for i := 0; i < count; i++ {
			s = append(s, newStruct())
		}
	}

	_ = s
}

func BenchmarkAllocStructLiteral(b *testing.B) {
	var s S

	for n := 0; n < b.N; n++ {
		s = S{
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

	_ = s
}

func BenchmarkAllocStructByFields(b *testing.B) {
	var s S

	for n := 0; n < b.N; n++ {
		var s S
		s.A = 1
		s.B = 2
		s.C = 3
		s.D = 4
		s.E = 5
		s.F = 6
		s.G = 7
		s.H = 8
		s.J = 9
		s.K = 10
	}

	_ = s
}
