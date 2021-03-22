package server

import (
	"log"
	"net"
	"net/http"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeoutMiddleware(t *testing.T) {
	wrappedHandler := TimeoutMiddleware(http.HandlerFunc(asyncSlowHandler), 5*time.Second)
	srv := NewServer(wrappedHandler)

	// start http server in background since it's blocking:
	go func() {
		log.Print("starting http server on ", srv.Addr)
		log.Print(srv.ListenAndServe())
	}()

	// workaround to start srv:
	runtime.Gosched()

	// client := &http.Client{
	// 	Timeout: 20 * time.Second,
	// }

	transport := &http.Transport{
		ResponseHeaderTimeout: 60 * time.Second,

		// defaults:
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "http://"+srv.Addr+"/foo", nil)
	require.NoError(t, err)

	resp, err := transport.RoundTrip(req)
	t.Logf("resp: %+v", resp)
	require.NoError(t, err)
	assert.Equal(t, http.StatusGatewayTimeout, resp.StatusCode)
}

func asyncSlowHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("got ", r.Method, r.URL, " request")

	select {
	case <-r.Context().Done():
		return

	// doing something slow:
	case <-time.After(10 * time.Second):
		break // from select
	}

	w.WriteHeader(http.StatusOK)
}
