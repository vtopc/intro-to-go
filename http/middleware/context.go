package main

import (
	"context"
	"net/http"
)

var userIDContextKey struct{}

func CtxWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}

// ReqWithUserID updated r with userID
func ReqWithUserID(r *http.Request, userID string) {
	ctx := CtxWithUserID(r.Context(), userID)

	req := r.Clone(ctx)
	*r = *req
}
