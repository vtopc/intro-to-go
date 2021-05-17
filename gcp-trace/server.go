package main

import (
	"net/http"
	"time"
)

func NewServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Handler: handler,

		Addr:              addr,
		ReadTimeout:       180 * time.Second,
		ReadHeaderTimeout: 180 * time.Second,
		WriteTimeout:      180 * time.Second,
	}
}
