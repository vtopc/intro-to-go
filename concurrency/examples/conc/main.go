package main

import (
	"context"
	"log"

	"concurrency"

	"github.com/sourcegraph/conc/pool"
)

func main() {
	requests := concurrency.GenerateRequests(concurrency.Count)

	DoAsync(context.TODO(), requests)
}

func DoAsync(ctx context.Context, requests [][]byte) {
	p := pool.NewWithResults[string]().WithMaxGoroutines(concurrency.TotalWorkers)

	for i, request := range requests {
		// https://github.com/golang/go/wiki/CommonMistakes/#using-goroutines-on-loop-iterator-variables
		id := i
		req := request

		log.Printf("sending request #%d", id)

		p.Go(func() string {
			return Work(ctx, id, req)
		})
	}

	res := p.Wait() // blocking

	getResults(res)
}

func Work(ctx context.Context, id int, req []byte) string {
	defer log.Printf("worker #%d: stopping\n", id)

	if ctx.Err() != nil {
		return ""
	}

	return concurrency.Md5sum(req)
}

func getResults(res []string) {
	log.Println("results:", res) // or write to DB...
}
