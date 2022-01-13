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

var (
	iface    interface{}
	i        int
	ifaceMap map[string]interface{}
	intMap   map[string]int
	st       S
	sl       Elems
)

//go:embed testdata/json.json
var jb []byte

func BenchmarkIfaceMapAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ifaceMap = newIfaceMap()
	}
}

func BenchmarkIntMapAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intMap = newIntMap()
	}
}

func BenchmarkSliceAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl = newSlice()
	}
}

func BenchmarkStructAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		st = newStruct()
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
		var ok bool
		iface, ok = v.(int)
		if !ok {
			panic(fmt.Sprintf("failed to assert %v(%T) into int", v, v))
		}
	}

}

func BenchmarkIntMapAvgSearch10Elems(b *testing.B) {
	m := newIntMap()

	for n := 0; n < b.N; n++ {
		i = m[avgKey]
	}
}

func BenchmarkIntSwitchAvgSearch10Elems(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i = switchSearch(avgKey)
	}
}

func BenchmarkIntSwitchWorstSearch10Elems(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i = switchSearch(worstKey)
	}
}

func BenchmarkSliceAvgSearch10Elems(b *testing.B) {
	s := newSlice()

	for n := 0; n < b.N; n++ {
		for _, e := range s {
			if e.K == avgKey {
				i = e.V
				break
			}
		}
	}
}

func BenchmarkSliceWorstSearch10Elems(b *testing.B) {
	s := newSlice()

	for n := 0; n < b.N; n++ {
		for _, e := range s {
			if e.K == worstKey {
				i = e.V
				break
			}
		}
	}
}

func BenchmarkStructSearch10Elems(b *testing.B) {
	s := newStruct()

	for n := 0; n < b.N; n++ {
		i = s.E
	}
}

func BenchmarkIfaceMapMarshal(b *testing.B) {
	v := newIfaceMap()

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkIntMapMarshal(b *testing.B) {
	v := newIntMap()

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStructMarshal(b *testing.B) {
	v := newStruct()

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
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
