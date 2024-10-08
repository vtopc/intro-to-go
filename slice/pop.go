package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	s2 := s
	for {
		if len(s)-1 == 0 {
			break
		}

		v := s[0] // get first element
		s = s[1:] // pop first element
		fmt.Printf("v: %v; s: %v; cap(%d); len(%d) \n", v, s, cap(s), len(s))
	}

	s[0] = 100500

	fmt.Println("s2: ", s2) // same underlying array
}
