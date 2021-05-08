package api

import (
	"context"
	"net/http"
)

func setMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenData, err := authenticateRequest(r)
		if err != nil {
			respondError(w, http.StatusUnauthorized, err)
			return
		}
		ctx := context.WithValue(r.Context(), "token", tokenData)
		r=r.WithContext(ctx)
		next(w, r)
	}
}
