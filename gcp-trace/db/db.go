package db

import (
	"context"
	"time"

	"gcp-trace/ltrace"

	"go.opencensus.io/trace"
)

// Query emulates some DB request
func Query(ctx context.Context, traceable bool) {
	if traceable {
		_, span := trace.StartSpan(ctx, ltrace.Prefix+"/db")
		defer span.End()
	}

	time.Sleep(50 * time.Millisecond)
}
