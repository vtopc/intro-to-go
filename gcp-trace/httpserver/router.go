package httpserver

import (
	"net/http"

	"gcp-trace/httpserver/healthz"
)

func NewRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/foo", http.HandlerFunc(Handler))
	router.Handle("/healthz", http.HandlerFunc(healthz.HealthCheck))

	return router
}
