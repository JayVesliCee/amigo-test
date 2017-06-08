package main

import (
	"context"
	"net/http"
)

type contextKey string

const configuration contextKey = "Configuration"

func contextMiddleware(service *service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), configuration, service.Config)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
