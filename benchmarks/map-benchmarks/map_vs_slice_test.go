package main

import (
	"strconv"
	"testing"
)

// prepare test data
var (
	sliceData []string
	mapData   map[string]int
)

// Global variable for Benchmarking - https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
var ok bool

func init() {
	const length = 2_000

	mapData = make(map[string]int, length)
	sliceData = make([]string, 0, length)
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		sliceData = append(sliceData, key)
		mapData[key] = i
	}
}

// linear search in slice
func sliceLookup(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func BenchmarkMapLookup2000ElemsAvg(b *testing.B) {
	key := "key1000"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, ok = mapData[key]
	}
}

func BenchmarkSliceLookup2000ElemsAvg(b *testing.B) {
	key := "key1000"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok = sliceLookup(sliceData, key)
	}
}

func BenchmarkSliceLookup2000ElemsWorst(b *testing.B) {
	key := "key9999" // worst case, not found
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok = sliceLookup(sliceData, key)
	}
}
