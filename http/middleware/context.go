package main

import (
	"context"
	"net/http"
)

var userIDContextKey struct{}

// ReqWithUserID updated r with userID
func ReqWithUserID(r *http.Request, userID string) {
	ctx := context.WithValue(r.Context(), userIDContextKey, userID)

	req := r.Clone(ctx)
	*r = *req
}
