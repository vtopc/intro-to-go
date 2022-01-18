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

func BenchmarkAllocMapIface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ifaceMap = newIfaceMap()
	}
}

func BenchmarkAllocMapInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intMap = newIntMap()
	}
}

func BenchmarkAllocSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl = newSlice()
	}
}

func BenchmarkAllocStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		st = newStruct()
	}
}

func BenchmarkWriteMapIface(b *testing.B) {
	m := make(map[string]interface{})

	for n := 0; n < b.N; n++ {
		m[avgKey] = 5
	}
}

func BenchmarkWriteMapInt(b *testing.B) {
	m := make(map[string]int)

	for n := 0; n < b.N; n++ {
		m[avgKey] = 5
	}
}

func BenchmarkWriteStruct(b *testing.B) {
	var s S

	for n := 0; n < b.N; n++ {
		s.E = 5
	}
}

func BenchmarkSearch10ElemsMapIfaceAvg(b *testing.B) {
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

func BenchmarkSearch10ElemsMapIntAvg(b *testing.B) {
	m := newIntMap()

	for n := 0; n < b.N; n++ {
		i = m[avgKey]
	}
}

func BenchmarkSearch10ElemsSwitchIntAvg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i = switchSearch(avgKey)
	}
}

func BenchmarkSearch10ElemsSwitchIntWorst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i = switchSearch(worstKey)
	}
}

func BenchmarkSearch10ElemsSliceAvg(b *testing.B) {
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

func BenchmarkSearch10ElemsSliceWorst(b *testing.B) {
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

func BenchmarkSearch10ElemsStruct(b *testing.B) {
	s := newStruct()

	for n := 0; n < b.N; n++ {
		i = s.E
	}
}

func BenchmarkMarshalMapIface(b *testing.B) {
	v := newIfaceMap()

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMarshalMapInt(b *testing.B) {
	v := newIntMap()

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMarshalStruct(b *testing.B) {
	v := newStruct()

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkUnmarshalMapInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v map[string]int
		err := json.Unmarshal(jb, &v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkUnmarshalMapIface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v map[string]interface{}
		err := json.Unmarshal(jb, &v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v S
		err := json.Unmarshal(jb, &v)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSetWriteIface(b *testing.B) {
	m := make(map[interface{}]struct{})

	for n := 0; n < b.N; n++ {
		m["E"] = struct{}{}
	}
}

func BenchmarkSetWriteString(b *testing.B) {
	m := make(map[string]struct{})

	for n := 0; n < b.N; n++ {
		m["E"] = struct{}{}
	}
}
