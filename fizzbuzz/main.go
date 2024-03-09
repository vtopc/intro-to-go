package main

import "fmt"

type T struct {
	v int
	s string
}

func fizzbuzz(i int) {
	m := []T{
		{v: 15, s: "fizzbuzz"},
		{v: 5, s: "buzz"},
		{v: 3, s: "fizz"},
	}

	for _, s := range m {
		if i%s.v == 0 {
			fmt.Printf("%s\n", s.s)
			return
		}
	}

	fmt.Printf("%d\n", i)
}

func main() {
	for i := 1; i < 100; i++ {
		fizzbuzz(i)
	}
}
