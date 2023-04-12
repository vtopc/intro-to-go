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
	totalWorkers := concurrency.TotalWorkers

	var resultsWG sync.WaitGroup

	// chan buffer should be tuned to the value when channels are not exhausted
	//  and workers are not waiting for the input:
	reqChan := make(chan []byte, totalWorkers)
	respChan := make(chan string, totalWorkers)

	go func() { // requests producer
		for i, request := range requests {
			log.Printf("sending request #%d", i)

			reqChan <- request
		}
		close(reqChan)
	}()

	resultsWG.Add(1)
	go getResults(&resultsWG, respChan)

	g, ctx := errgroup.WithContext(ctx) // use `errgroup.Group` literal if you don't need to cancel context on the first error
	g.SetLimit(totalWorkers)

	go respChanCloser(ctx, respChan)

	// consumer loop
	var i int

loop:
	for {
		select {
		case <-ctx.Done():
			break loop

		case req, ok := <-reqChan:
			if !ok {
				break loop
			}

			id := i
			g.Go(func() (err error) {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("recovered panic: %s", r)
					}
				}()

				// if i == 42 {
				// 	return errors.New("error-42")
				// }

				Work(ctx, id, req, respChan)

				return nil
			})

			i++
		}
	}

	// kind a shut-down

	if err := g.Wait(); err != nil { // blocking
		fmt.Println("worker error:", err)
	}

	resultsWG.Wait() // blocking
}

func Work(ctx context.Context, id int, req []byte, respChan chan<- string) {
	// if id == 42 {
	// 	panic("42")
	// }

	s := concurrency.Md5sum(req)

	select {
	case respChan <- s:
		log.Printf("worker #%d: send: %s\n", id, s)

	case <-ctx.Done():
		log.Printf("worker #%d: cancelling\n", id)
		return
	}
}

func respChanCloser(ctx context.Context, respChan chan<- string) {
	<-ctx.Done() // goroutine is blocked

	log.Println("all workers are done, closing respChan")

	close(respChan) // sending a "signal" to func getResults(), that there will be no more messages.
}

func getResults(wg *sync.WaitGroup, respChan <-chan string) {
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
