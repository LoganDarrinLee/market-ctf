package middleware

import (
	"context"
	"net/http"

	"github.com/LoganDarrinLee/market-ctf/internal/common"
)

// Add request context to the handler.
// Can be easily midified to add more context values. Documented below.
func WithRequestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := common.GenerateNewID()

		// New context value
		ctx := context.WithValue(r.Context(), common.RequestIDKey, requestID)
		// Add more context here with the following.
		// ctx = context.WithValue(ctx, <key>, <value>)

		// Pass the new context along the handler chain
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
