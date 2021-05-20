package main

import (
	"context"
	"errors"
	"net/http"
	"runtime"
	"testing"
	"time"

	"gcp-trace/httpserver"
	"gcp-trace/ltrace"

	"github.com/stretchr/testify/require"
)

const addr = "127.0.0.1:8081"

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func BenchmarkWithoutTrace(b *testing.B) {
	router := httpserver.NewRouter(false)
	srv := httpserver.NewServer(addr, router)
	go func() {
		b.Logf("starting http server on %s", addr)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			b.Logf("http srv error: %s", err)
		}
	}()
	defer func() {
		_ = srv.Shutdown(context.Background())
	}()

	// workaround for srv to start:
	runtime.Gosched()

	req, err := http.NewRequest(http.MethodGet, "http://"+addr+"/foo", nil)
	require.NoError(b, err)

	for n := 0; n < b.N; n++ {
		_, err := client.Do(req)
		require.NoError(b, err)
	}
}

func BenchmarkWithTrace(b *testing.B) {
	ltrace.RegisterTrace()

	router := ltrace.HTTPHandlerWrapper(httpserver.NewRouter(true))
	srv := httpserver.NewServer(addr, router)
	go func() {
		b.Logf("starting http server on %s", addr)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			b.Logf("http srv error: %s", err)
		}
	}()
	defer func() {
		_ = srv.Shutdown(context.Background())
	}()

	// workaround for srv to start:
	runtime.Gosched()

	req, err := http.NewRequest(http.MethodGet, "http://"+addr+"/foo", nil)
	require.NoError(b, err)

	for n := 0; n < b.N; n++ {
		_, err := client.Do(req)
		require.NoError(b, err)
	}
}
