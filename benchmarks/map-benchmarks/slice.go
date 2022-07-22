package main

type Elems []Elem

type PtrElems []*Elem

type Elem struct {
	K string
	V int
}

type HugeElem [1024]byte

type HugeElems []HugeElem

type HugePtrElems []*HugeElem

func newSlice() Elems {
	return Elems{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"D", 4},
		{"E", 5},
		{"F", 6},
		{"G", 7},
		{"H", 8},
		{"J", 9},
		{"K", 10},
	}
}

func newPtrSlice() PtrElems {
	return PtrElems{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"D", 4},
		{"E", 5},
		{"F", 6},
		{"G", 7},
		{"H", 8},
		{"J", 9},
		{"K", 10},
	}
}

func newHugeSlice() HugeElems {
	return make(HugeElems, 10)
}

func newHugePtrSlice() HugePtrElems {
	return HugePtrElems{
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
		&HugeElem{},
	}
}
