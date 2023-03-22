package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"concurrency"

	"golang.org/x/sync/errgroup"
)

func main() {
	requests := concurrency.GenerateRequests(concurrency.Count)

	DoAsync(context.TODO(), requests)
}

func DoAsync(ctx context.Context, requests [][]byte) {
	// chan buffer should be tuned to the value when channels are not exhausted
	//  and workers are not waiting for the input:
	respChan := make(chan string, concurrency.TotalWorkers)

	var resultsWG sync.WaitGroup
	resultsWG.Add(1)
	go getResults(respChan, &resultsWG)

	group, _ := errgroup.WithContext(ctx) // use `errgroup.Group` literal if you don't need to cancel on first error
	group.SetLimit(concurrency.TotalWorkers)

	for i, request := range requests {
		// https://github.com/golang/go/wiki/CommonMistakes/#using-goroutines-on-loop-iterator-variables
		id := i
		req := request

		log.Printf("sending request #%d", id)

		group.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered panic: %s", r)
				}
			}()

			Work(ctx, id, req, respChan)

			return nil
		})
	}

	err := group.Wait() // blocking
	if err != nil {
		fmt.Println("worker error:", err)
	}

	close(respChan)

	resultsWG.Wait() // blocking
}

func Work(ctx context.Context, id int, req []byte, respChan chan<- string) {
	s := concurrency.Md5sum(req)

	select {
	case respChan <- s:
		log.Printf("worker #%d: send: %s\n", id, s)

	case <-ctx.Done():
		log.Printf("worker #%d: stopping\n", id)
		return
	}
}

func getResults(respChan <-chan string, wg *sync.WaitGroup) {
	batchSize := concurrency.Count
	res := make([]string, 0, batchSize)

	for {
		s, ok := <-respChan
		if !ok {
			break
		}

		log.Println("getResults: got from workers:", s)

		res = append(res, s)

		if len(res) == batchSize {
			log.Println("results:", res) // or write to DB...

			res = make([]string, 0, batchSize)
		}
	}

	if len(res) != 0 {
		log.Println("final results:", res)
	}

	// all results are saved

	wg.Done()
}
