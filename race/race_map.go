package main

import (
	"fmt"
	"log"
	"time"
)

type T map[string]int

func a(x T) {
	x["a"] = 1
}

func b(x T) {
	x["b"] = 2
}

func race(x T) {
	go a(x)
	go b(x)
}

// data race in not recovered
func raceWithRecover(x T) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic in a: %s", r)
			}
		}()

		a(x)
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic in b: %s", r)
			}
		}()

		b(x)
	}()
}

func main() {
	x := make(T)
	race(x)
	time.Sleep(time.Nanosecond)
	fmt.Println(x)
}
