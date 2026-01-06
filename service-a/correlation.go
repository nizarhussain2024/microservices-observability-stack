package main

import (
	"context"
	"net/http"
)

type correlationKey string

const correlationIDKey correlationKey = "correlation_id"

func getCorrelationID(r *http.Request) string {
	// Check header first
	if id := r.Header.Get("X-Correlation-ID"); id != "" {
		return id
	}
	// Generate if not present
	return generateRequestID()
}

func withCorrelationID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, correlationIDKey, id)
}

func correlationIDFromContext(ctx context.Context) string {
	if id, ok := ctx.Value(correlationIDKey).(string); ok {
		return id
	}
	return ""
}

func correlationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		correlationID := getCorrelationID(r)
		w.Header().Set("X-Correlation-ID", correlationID)
		
		ctx := withCorrelationID(r.Context(), correlationID)
		next(w, r.WithContext(ctx))
	}
}



