package main

import (
	"math/rand"
	"testing"
	"time"
)

var (
	global int
	input  = getRandomElements()
)

func Benchmark_sumElements(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumElements(input)
	}
	global = local
}

func Benchmark_sumByIndex(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumByIndex(input)
	}
	global = local
}

func getRandomElements() []S {
	n := 1_000
	res := make([]S, n)
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	for i := 0; i < n; i++ {
		res[i].i = rnd.Int()
	}
	return res
}
