package main

import (
	"fmt"
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

func main() {
	race()
	time.Sleep(time.Nanosecond)
	fmt.Println(x)
}
