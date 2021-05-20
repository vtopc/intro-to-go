package httpserver

import (
	"net/http"
	"strconv"

	"gcp-trace/db"
	"gcp-trace/ltrace"

	"go.opencensus.io/trace"
)

type Handler struct {
	Traceable bool
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // default context
	if h.Traceable {
		// get current span:
		//  span := trace.FromContext(r.Context())
		// or create new span:
		var span *trace.Span
		ctx, span = trace.StartSpan(r.Context(), ltrace.Prefix+"/handler")
		defer span.End()

		w.Header().Set("X-B3-TraceId", span.SpanContext().TraceID.String())
		w.Header().Set("X-B3-Sampled", strconv.Itoa(btoi(span.SpanContext().IsSampled())))

		span.AddAttributes(trace.Int64Attribute("custom", 42))
	}

	db.Query(ctx, h.Traceable)

	// uncomment for client propagation:
	// // The trace ID from the incoming request will be
	// // propagated to the outgoing request.
	// req, err := http.NewRequestWithContext(
	// 	ctx,
	// 	http.MethodGet,
	// 	"https://www.google.com",
	// 	nil,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// // The outgoing request will be traced with r's trace ID.
	// resp, err := tracedClient.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// // Because we don't read the resp.Body, need to manually call Close().
	// _ = resp.Body.Close()

	_, _ = w.Write([]byte("bar"))
}

func btoi(b bool) int {
	if b {
		return 1
	}

	return 0
}
