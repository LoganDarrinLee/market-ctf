package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/LoganDarrinLee/market-ctf/internal/common"
	"github.com/LoganDarrinLee/market-ctf/internal/routing"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request context: ", r.Context().Value(common.RequestIDKey))

		next.ServeHTTP(w, r)
	})
}

func VendorPage(h *routing.Handler, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sessionTokenPermission(r.Context(), "vendor") {
			next.ServeHTTP(w, r)
		} else {

		}
	})
}

func sessionTokenPermission(ctx context.Context, requiredAccess string) bool {
	return false
}
