package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"log"
	"sync"
)

const (
	Count        = 50
	TotalWorkers = 5
)

func main() {
	requests := generateRequest(Count)

	DoAsync(requests)

	// DoSync(requests)
}

func DoAsync(requests [][]byte) {
	var wg sync.WaitGroup

	// chan buffer should be tuned to the value when channels are not exhausted
	//  and workers are not waiting for the input:
	reqChan := make(chan []byte, TotalWorkers)
	respChan := make(chan string, TotalWorkers)

	doneChan := make(chan struct{}, 1)

	go func() { // size of reqChan is only 5
		for i, request := range requests {
			log.Printf("sending request #%d", i)

			reqChan <- request
		}
		close(reqChan)
	}()

	wg.Add(TotalWorkers)
	for id := 1; id <= TotalWorkers; id++ {
		// starting workers
		go Work(id, &wg, reqChan, respChan)
	}

	go resChanCloser(&wg, respChan)

	go getResults(respChan, doneChan)

	<-doneChan // blocking
}

func Work(id int, wg *sync.WaitGroup, reqChan <-chan []byte, respChan chan<- string) {
	log.Printf("worker #%d: started\n", id)

	for data := range reqChan {
		s := md5sum(data)

		log.Printf("worker #%d: sending: %s\n", id, s)

		respChan <- s
	}

	log.Printf("worker #%d: done\n", id)
	wg.Done()
}

func resChanCloser(wg *sync.WaitGroup, respChan chan<- string) {
	wg.Wait() // goroutine is blocked

	log.Println("all workers are done, closing respChan")

	close(respChan) // sending a "signal" to func getResults(), that there will be no more messages.
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
