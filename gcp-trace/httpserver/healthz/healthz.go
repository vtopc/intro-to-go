package healthz

import (
	"net/http"
	"strings"
)

const Path = "/healthz"

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("OK"))
}

func IsHealthEndpoint(r *http.Request) bool {
	if strings.HasSuffix(r.URL.String(), Path) {
		return true
	}

	return false
}
