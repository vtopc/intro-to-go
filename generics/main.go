package main

import (
	"fmt"
	"time"
)

type Duration interface {
	Seconds() float64
}

type Durations[T Duration] []T

func interfaces(dd []Duration) {
	fmt.Println(dd)
}

func generics[T Duration](dd Durations[T]) {
	fmt.Println(dd)
}

func main() {
	durations := []time.Duration{time.Second, time.Minute, time.Hour}
	// interfaces(durations) // <-- compile-time error. Need to repack into []Duration.

	generics(durations)
}
