package server

import (
	"context"
	"errors"
	"net/http"
	"time"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Handler: handler,

		Addr:              "127.0.0.1:8080",
		ReadTimeout:       180 * time.Second,
		ReadHeaderTimeout: 180 * time.Second,
		WriteTimeout:      180 * time.Second,
	}
}

// TimeoutMiddleware is kinda http.TimeoutHandler, but returns 504 instead 503.
// TODO: rewrite to http.TimeoutHandler?
func TimeoutMiddleware(h http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout) // context.DeadlineExceeded
			defer cancel()
			r = r.WithContext(ctx)

			done := make(chan struct{})

			go func() {
				h.ServeHTTP(w, r) // call original
				close(done)
			}()

			select {
			case <-ctx.Done():
				err := ctx.Err()
				if errors.Is(err, context.DeadlineExceeded) ||
					errors.Is(err, context.Canceled) {

					w.WriteHeader(http.StatusGatewayTimeout)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}

				_, _ = w.Write([]byte(err.Error()))

			case <-done:
				// all good
				return
			}
		},
	)
}
