package main

import (
	"context"
	"log"
	"sync"

	"concurrency"
)

func main() {
	requests := concurrency.GenerateRequests(concurrency.Count)

	DoAsync(context.TODO(), requests)
}

func DoAsync(ctx context.Context, requests [][]byte) {
	totalWorkers := concurrency.TotalWorkers

	var workersWG, resultsWG sync.WaitGroup

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

	workersWG.Add(totalWorkers)
	for id := 1; id <= totalWorkers; id++ {
		// starting workers
		go Work(ctx, id, &workersWG, reqChan, respChan)
	}

	go respChanCloser(&workersWG, respChan)

	resultsWG.Wait() // blocking
}

func Work(ctx context.Context, id int, wg *sync.WaitGroup, reqChan <-chan []byte, respChan chan<- string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered panic: %s", r)
		}
	}()

	log.Printf("worker #%d: started\n", id)

	for data := range reqChan {
		s := concurrency.Md5sum(data)

		select {
		case respChan <- s:
			log.Printf("worker #%d: send: %s\n", id, s)

		case <-ctx.Done():
			log.Printf("worker #%d: stopping\n", id)
			return
		}
	}

	log.Printf("worker #%d: done\n", id)
	wg.Done()
}

func respChanCloser(wg *sync.WaitGroup, respChan chan<- string) {
	wg.Wait() // goroutine is blocked

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
