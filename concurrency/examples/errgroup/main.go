package main

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	Count        = 50
	TotalWorkers = 5
)

func main() {
	requests := generateRequest(Count)

	DoAsync(context.TODO(), requests)

	// DoSync(requests)
}

func DoAsync(ctx context.Context, requests [][]byte) {
	// chan buffer should be tuned to the value when channels are not exhausted
	//  and workers are not waiting for the input:
	reqChan := make(chan []byte, TotalWorkers)
	respChan := make(chan string, TotalWorkers)

	var resultsWG sync.WaitGroup
	resultsWG.Add(1)
	go getResults(respChan, &resultsWG)

	g, _ := errgroup.WithContext(ctx) // use `errgroup.Group` literal if you don't need to cancel on first error
	g.SetLimit(TotalWorkers)

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

			Work(ctx, id, req, respChan)

			return nil
		})
	}

	close(reqChan)

	err := g.Wait() // blocking
	if err != nil {
		fmt.Println("worker error:", err)
	}

	close(respChan)

	resultsWG.Wait() // blocking
}

func Work(ctx context.Context, id int, req []byte, respChan chan<- string) {
	s := md5sum(req)

	select {
	case respChan <- s:
		log.Printf("worker #%d: send: %s\n", id, s)

	case <-ctx.Done():
		log.Printf("worker #%d: stopping\n", id)
		return
	}
}

func getResults(respChan <-chan string, wg *sync.WaitGroup) {
	batchSize := Count
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

func generateRequest(length int) [][]byte {
	res := make([][]byte, 0, length)

	for i := 1; i <= length; i++ {
		b := make([]byte, 2*1024*1024) // 2 Mb

		_, err := rand.Read(b)
		if err != nil {
			panic(fmt.Errorf("random failed: %s", err))
		}

		res = append(res, b)
	}

	log.Println("generated requests")

	return res
}

func md5sum(data []byte) string {
	// time.Sleep(100 * time.Millisecond)

	return fmt.Sprintf("%x", md5.Sum(data))
}

func DoSync(requests [][]byte) {
	res := make([]string, 0, Count)

	for i, request := range requests {
		log.Printf("sending request #%d\n", i)

		s := md5sum(request)

		log.Printf("saving: %s\n", s)

		res = append(res, s)
		log.Printf("saved: %s\n", s)
	}

	log.Println("results:", res)
}
