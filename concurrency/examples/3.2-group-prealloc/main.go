package main

import (
	"context"
	"fmt"
	"log"

	"concurrency"

	"golang.org/x/sync/errgroup"
)

func main() {
	requests := concurrency.GenerateRequests(concurrency.Count)

	DoAsync(context.TODO(), requests)
}

func DoAsync(ctx context.Context, requests [][]byte) {
	// https://stackoverflow.com/questions/49879322/can-i-concurrently-write-different-slice-elements
	resp := make([]string, concurrency.Count)

	g, _ := errgroup.WithContext(ctx) // use `errgroup.Group` literal if you don't need to cancel context on the first error
	g.SetLimit(concurrency.TotalWorkers)

	for i, request := range requests {
		// https://github.com/golang/go/wiki/CommonMistakes/#using-goroutines-on-loop-iterator-variables
		id := i
		req := request

		log.Printf("sending request #%d", id)

		g.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered panic: %s", r)
				}
			}()

			resp[id], err = Work(ctx, id, req)
			if err != nil {
				return err
			}

			return nil
		})
	}

	err := g.Wait() // blocking
	if err != nil {
		fmt.Println("worker error:", err)
	}

	getResults(resp)
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
