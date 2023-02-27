package main

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"log"

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

	doneChan := make(chan struct{}, 1)

	go getResults(respChan, doneChan)

	g, _ := errgroup.WithContext(ctx)
	g.SetLimit(TotalWorkers)

	for i, request := range requests {
		id := i
		req := request

		log.Printf("sending request #%d", id)

		g.Go(func() error {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered panic:", r)
				}
			}()

			Work(id, req, respChan)

			return nil
		})
	}

	close(reqChan)

	err := g.Wait() // blocking
	if err != nil {
		fmt.Println("worker error:", err)
	}

	close(respChan)

	<-doneChan // blocking
}

func Work(id int, req []byte, respChan chan<- string) {
	s := md5sum(req)

	log.Printf("sending response #%d: %s\n", id, s)

	respChan <- s
}

func getResults(respChan <-chan string, doneChan chan<- struct{}) {
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

	close(doneChan)
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
