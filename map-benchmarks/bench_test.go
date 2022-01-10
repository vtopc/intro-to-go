package main

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed testdata/json.json
var jb []byte

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

func BenchmarkIfaceMapMarshal(b *testing.B) {
	v := newIfaceMap()

	for n := 0; n < b.N; n++ {
		_, _ = json.Marshal(v)
	}
}

func BenchmarkIntMapMarshal(b *testing.B) {
	v := newIntMap()

	for n := 0; n < b.N; n++ {
		_, _ = json.Marshal(v)
	}
}

func BenchmarkStructMarshal(b *testing.B) {
	v := newStruct()

	for n := 0; n < b.N; n++ {
		_, _ = json.Marshal(v)
	}
}

func BenchmarkIntMapUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v map[string]int
		err := json.Unmarshal(jb, &v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkIfaceMapUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v map[string]interface{}
		err := json.Unmarshal(jb, &v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSliceUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v S
		err := json.Unmarshal(jb, &v)
		if err != nil {
			panic(err)
		}
	}
}
