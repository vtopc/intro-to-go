package server

import (
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
