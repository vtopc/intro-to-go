package main

import (
	"math/rand/v2"
	"testing"
)

const n = 1_000

var (
	global int
	inputV = getRandomVElements()
	inputP = getRandomPElements()
)

func Benchmark_sumValueElements(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumValueElements(inputV)
	}
	global = local
}

func Benchmark_sumValuesByIndex(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumValuesByIndex(inputV)
	}
	global = local
}

func Benchmark_sumPointerElements(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumPointerElements(inputP)
	}
	global = local
}

func Benchmark_sumPointersByIndex(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumPointersByIndex(inputP)
	}
	global = local
}

// TODO: generics?
func getRandomVElements() []S {
	ret := make([]S, n)
	for i := 0; i < n; i++ {
		ret[i].i = rand.Int()
	}

	return ret
}

func getRandomPElements() []*S {
	ret := make([]*S, 0, n)
	for i := 0; i < n; i++ {
		e := &S{i: rand.Int()}
		ret = append(ret, e)
	}

	return ret
}
