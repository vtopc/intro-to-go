package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	for {
		if len(s) == 0 {
			break
		}

		v := s[0] // get first element
		s = s[1:] // pop first element
		fmt.Printf("v: %v; s: %v; cap(%d); len(%d) \n", v, s, cap(s), len(s))
	}
}
