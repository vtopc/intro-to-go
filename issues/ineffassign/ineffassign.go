package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fn1()
	err = fn2()
	if err != nil {
		fmt.Println(err)
	}
}

func fn1() error {
	return nil
}

func fn2() error {
	return errors.New("some error")
}
