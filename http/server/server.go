package server

import (
	"net/http"
	"time"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Handler: handler,

		// TODO: move to configs:
		Addr: "127.0.0.1:8080",
		// ReadTimeout: 1 * time.Second,
		// ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}
