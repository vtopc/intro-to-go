package server

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestServerTimeout(t *testing.T) {
	srv := NewServer(http.HandlerFunc(slowHandler))

	// test http.Server.WriteTimeout:
	srv.WriteTimeout = 1 * time.Second

	// start http server in background since it's blocking:
	go func() {
		log.Print("starting http server on ", srv.Addr)
		log.Print(srv.ListenAndServe())
	}()

	// workaround for srv to start:
	runtime.Gosched()

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://"+srv.Addr+"/foo", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	t.Logf("resp: %+v", resp)
	require.EqualError(t, err, `Get "http://127.0.0.1:8080/foo": EOF`)

	// assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("got ", r.Method, r.URL, " request")

	time.Sleep(10 * time.Second)

	w.WriteHeader(http.StatusOK)
}
