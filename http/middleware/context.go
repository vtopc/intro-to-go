package main

import (
	"context"
	"net/http"
)

var userIDContextKey struct{}

func CtxWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}

// ReqWithCtx updates provided http.Request with Context
func ReqWithCtx(ctx context.Context, r *http.Request) {
	// To change the context of a request, such as an incoming request you
	// want to modify before sending back out, use Request.Clone.
	// https://pkg.go.dev/net/http@go1.19.6#Request.WithContext
	req := r.Clone(ctx)

	*r = *req
}
