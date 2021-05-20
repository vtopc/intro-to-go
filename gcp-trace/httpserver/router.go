package httpserver

import (
	"net/http"

	"gcp-trace/httpserver/healthz"
)

func NewRouter(traceable bool) http.Handler {
	handler := Handler{Traceable: traceable}

	router := http.NewServeMux()
	router.Handle("/foo", handler)
	router.Handle("/healthz", http.HandlerFunc(healthz.HealthCheck))

	return router
}
