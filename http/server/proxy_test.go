package server

import (
	"log"
	"net"
	"net/http"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTransportResponseHeaderTimeout(t *testing.T) {
	srv := NewServer(http.HandlerFunc(slowHandler))
	defer srv.Close()

	// start http server in background since it's blocking:
	go func() {
		log.Print("starting http server on ", srv.Addr)
		log.Print(srv.ListenAndServe())
	}()

	// workaround for srv to start:
	runtime.Gosched()

	transport := &http.Transport{
		// test http.Transport.ResponseHeaderTimeout:
		ResponseHeaderTimeout: 5 * time.Second,

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
	require.EqualError(t, err, `net/http: timeout awaiting response headers`)

	// assert.Equal(t, http.StatusOK, resp.StatusCode)
}
