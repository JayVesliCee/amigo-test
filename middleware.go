package main

import (
	"context"
	"net/http"
)

type contextKey string

const (
	configuration contextKey = "Configuration"
	psqlDB        contextKey = "psqlDB"
)

func contextMiddleware(service *service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), configuration, service.Config)
			ctx = context.WithValue(ctx, psqlDB, service.DB)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
