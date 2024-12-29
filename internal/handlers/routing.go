package handlers

import (
	"net/http"
)

func BaseRoutes(router *http.ServeMux, h *Handler) {

	// Get methods
	router.HandleFunc("GET /", h.indexHandler)
	router.HandleFunc("GET /about", h.aboutHandler)
}
