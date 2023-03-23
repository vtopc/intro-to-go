package concurrency

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"runtime"
)

const Count = 100

var TotalWorkers = runtime.NumCPU()

func GenerateRequests(length int) [][]byte {
	res := make([][]byte, 0, length)

	for i := 1; i <= length; i++ {
		b := make([]byte, 2*1024*1024) // 2 Mb

		_, err := rand.Read(b)
		if err != nil {
			panic(fmt.Errorf("random failed: %s", err))
		}

		res = append(res, b)
	}

	return res
}

func Md5sum(data []byte) string {
	// time.Sleep(100 * time.Millisecond)

	return fmt.Sprintf("%x", md5.Sum(data))
}
