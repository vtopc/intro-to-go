package main

type Elems []Elem

type Elem struct {
	K string
	V int
}

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
