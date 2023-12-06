package main

import (
	"fmt"
	"log"
	"time"
)

var x = make(map[string]int) // shared(package) variable

func a() {
	x["a"] = 1
}

func b() {
	x["b"] = 2
}

func race() {
	go a()
	go b()
}

// data race in not recovered
func raceWithRecover() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic in a: %s", r)
			}
		}()

		a()
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic in b: %s", r)
			}
		}()

		b()
	}()
}

func main() {
	race()
	time.Sleep(time.Nanosecond)
	fmt.Println(x)
}
