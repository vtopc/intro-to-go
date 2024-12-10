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

func Benchmark_sumValueElements(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumValueElements(input)
	}
	global = local
}

func Benchmark_sumValuesByIndex(b *testing.B) {
	var local int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = sumValuesByIndex(input)
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
