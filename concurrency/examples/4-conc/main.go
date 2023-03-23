package main

import (
	"context"
	"fmt"
	"log"

	"concurrency"

	"github.com/sourcegraph/conc/panics"
	"github.com/sourcegraph/conc/pool"
)

func main() {
	requests := concurrency.GenerateRequests(concurrency.Count)

	DoAsync(context.TODO(), requests)
}

func DoAsync(ctx context.Context, requests [][]byte) {
	p := pool.NewWithResults[string]().
		WithContext(ctx).
		WithMaxGoroutines(concurrency.TotalWorkers)

	for i, request := range requests {
		// https://github.com/golang/go/wiki/CommonMistakes/#using-goroutines-on-loop-iterator-variables
		id := i
		req := request

		log.Printf("sending request #%d", id)

		p.Go(func(c context.Context) (string, error) {
			var err error
			var ret string
			r := panics.Try(func() {
				ret, err = Work(c, id, req)
			})
			if r != nil {
				err = fmt.Errorf("recovered panic: %w", r.AsError())
			}

			return ret, err
		})
	}

	res, err := p.Wait() // blocking
	if err != nil {
		log.Printf("got error: %s", err)
		return
	}

	getResults(res)
}

func Work(ctx context.Context, id int, req []byte) (string, error) {
	defer log.Printf("worker #%d: stopping\n", id)

	// if id == 42 {
	// 	panic("42")
	// }

	err := ctx.Err()
	if err != nil {
		return "", err
	}

	return concurrency.Md5sum(req), nil
}

func getResults(res []string) {
	log.Println("results:", res) // or write to DB...
}
