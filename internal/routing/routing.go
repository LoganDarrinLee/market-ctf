package routing

import (
	"net/http"
)

// Register routes from the handler with the mux router.
func ConfigureRoutes(router *http.ServeMux, h *Handler) {

	// Get methods
	// router.Handle("GET /test", middleware.CheckAuth(h.authHandler()))

	// Standard landing pages of the market. Auth not required.
	router.HandleFunc("GET /", h.indexHandler)
	router.HandleFunc("GET /about", h.aboutHandler)
}
