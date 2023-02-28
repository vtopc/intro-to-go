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
	req := r.Clone(ctx)
	*r = *req
}
