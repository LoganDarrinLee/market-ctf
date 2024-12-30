package routing

import (
	"net/http"

	"github.com/LoganDarrinLee/market-ctf/internal/middleware"
)

func ConfigureRoutes(router *http.ServeMux, h *Handler) {

	// Get methods
	router.Handle("GET /test", middleware.CheckAuth(h.authHandler()))

	router.HandleFunc("GET /", h.indexHandler)
	router.HandleFunc("GET /about", h.aboutHandler)
}
