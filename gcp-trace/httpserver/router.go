package httpserver

import (
	"net/http"

	"gcp-trace/httpserver/healthz"
)

func NewRouter(trace bool) http.Handler {
	handler := Handler{Trace: trace}

	router := http.NewServeMux()
	router.Handle("/foo", handler)
	router.Handle("/healthz", http.HandlerFunc(healthz.HealthCheck))

	return router
}
