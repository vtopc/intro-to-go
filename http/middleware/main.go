package main

import (
	"log"
	"net/http"
	"time"
)

// Middleware function that logs the time taken for each request
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		userID := (r.Context().Value(userIDContextKey)).(string)

		log.Printf("%s %s(latency=%s; user_id=%s)", r.Method, r.URL.Path, time.Since(start), userID)
	}
}

// Handler function that returns a simple message
func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := CtxWithUserID(r.Context(), "42")

	ReqWithCtx(ctx, r)

	_, _ = w.Write([]byte("Hello, World!"))
}

func main() {
	// Create a new HTTP router
	router := http.NewServeMux()

	// Register the logging middleware for all routes
	router.HandleFunc("/hello", loggingMiddleware(helloHandler))

	// Start the server on port 8080
	log.Printf("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
