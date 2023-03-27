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

	// chan buffer should be tuned to the value when channels are not exhausted
	//  and workers are not waiting for the input:
	respChan := make(chan string, totalWorkers)

	var resultsWG sync.WaitGroup
	resultsWG.Add(1)
	go getResults(respChan, &resultsWG)

	g, ctx := errgroup.WithContext(ctx) // use `errgroup.Group` literal if you don't need to cancel context on the first error
	g.SetLimit(totalWorkers)

	for i, req := range requests {
		i, req := i, req // https://github.com/golang/go/wiki/CommonMistakes/#using-goroutines-on-loop-iterator-variables

		log.Printf("sending request #%d", i)

		g.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered panic: %s", r)
				}
			}()

			// if i == 42 {
			// 	return errors.New("error-42")
			// }

			Work(ctx, i, req, respChan)

			return nil
		})
	}

	err := g.Wait() // blocking
	if err != nil {
		fmt.Println("worker error:", err)
	}

	close(respChan)

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
