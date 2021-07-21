package main

import (
	"fmt"
)

func captured() (i int) {
	i = 1

	defer func(j int) {
		fmt.Println("captured deferred:", j)
	}(i)

	i++

	return
}

func latest() (i int) {
	i = 1

	defer func() {
		fmt.Println("latest deferred:", i)
	}()

	i++

	return
}

func main() {
	fmt.Println("captured:", captured())
	fmt.Println("latest:", latest())
}
