package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", http.NoBody)
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Millisecond)

	_, err = client.Do(req)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("context.DeadlineExceeded")
	}
}
