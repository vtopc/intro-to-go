package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
)

const (
	avgKey   = "E"
	worstKey = "K"
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

func BenchmarkSliceAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newSlice()
	}
}

func BenchmarkStructAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newStruct()
	}
}

func BenchmarkIfaceMapWrite(b *testing.B) {
	m := make(map[string]interface{})

	for n := 0; n < b.N; n++ {
		m[avgKey] = 5
	}
}

func BenchmarkIntMapWrite(b *testing.B) {
	m := make(map[string]int)

	for n := 0; n < b.N; n++ {
		m[avgKey] = 5
	}
}

func BenchmarkStructWrite(b *testing.B) {
	var s S

	for n := 0; n < b.N; n++ {
		s.E = 5
	}
}

func BenchmarkIfaceMapAvgSearch10Elems(b *testing.B) {
	m := newIfaceMap()

	for n := 0; n < b.N; n++ {
		v := m[avgKey]
		_, ok := v.(int)
		if !ok {
			panic(fmt.Sprintf("failed to assert %v(%T) into int", v, v))
		}
	}
}

func BenchmarkIntMapAvgSearch10Elems(b *testing.B) {
	m := newIntMap()

	for n := 0; n < b.N; n++ {
		_ = m[avgKey]
	}
}

func BenchmarkIntSwitchAvgSearch10Elems(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = switchSearch(avgKey)
	}
}

func BenchmarkIntSwitchWorstSearch10Elems(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = switchSearch("K")
	}
}

func BenchmarkStructSearch10Elems(b *testing.B) {
	s := newStruct()

	for n := 0; n < b.N; n++ {
		_ = s.E
	}
}

func BenchmarkSliceAvgSearch10Elems(b *testing.B) {
	s := newSlice()

	for n := 0; n < b.N; n++ {
		for _, e := range s {
			if e.K == avgKey {
				_ = e.V
				break
			}
		}
	}
}

func BenchmarkSliceWorstSearch10Elems(b *testing.B) {
	s := newSlice()

	for n := 0; n < b.N; n++ {
		for _, e := range s {
			if e.K == "K" {
				_ = e.V
				break
			}
		}
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

func BenchmarkStructUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v S
		err := json.Unmarshal(jb, &v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkIfaceSetWrite(b *testing.B) {
	m := make(map[interface{}]struct{})

	for n := 0; n < b.N; n++ {
		m["E"] = struct{}{}
	}
}

func BenchmarkStringSetWrite(b *testing.B) {
	m := make(map[string]struct{})

	for n := 0; n < b.N; n++ {
		m["E"] = struct{}{}
	}
}
