package main

import (
	"fmt"
)

func captured() (i int) {
	i = 1

	defer func(j int) {
		fmt.Println("captured defer:", j)
	}(i)

	i++

	return
}

func pointer() (i int) {
	i = 1

	defer func(j *int) {
		fmt.Println("pointer defer:", *j)
	}(&i)

	i++

	return
}

func latest() (i int) {
	i = 1

	defer func() {
		fmt.Println("latest defer:", i)
	}()

	i++

	return
}

func main() {
	fmt.Println("captured:", captured())
	fmt.Println("pointer:", pointer())
	fmt.Println("latest:", latest())
}
