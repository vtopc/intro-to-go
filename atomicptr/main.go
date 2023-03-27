package main

import (
	"sync/atomic"
)

type S struct {
	i atomic.Pointer[int]
}

func main() {}
