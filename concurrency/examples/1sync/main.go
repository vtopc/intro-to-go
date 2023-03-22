package main

import (
	"log"

	"concurrency"
)

func main() {
	requests := concurrency.GenerateRequests(concurrency.Count)

	DoSync(requests)
}

func DoSync(requests [][]byte) {
	res := make([]string, 0, concurrency.Count)

	for i, request := range requests {
		log.Printf("sending request #%d\n", i)

		s := concurrency.Md5sum(request)

		log.Printf("saving: %s\n", s)

		res = append(res, s)
		log.Printf("saved: %s\n", s)
	}

	log.Println("results:", res)
}
