// Routing configuration for the main mux router
package routing

import (
	"net/http"

	"github.com/LoganDarrinLee/market-ctf/internal/handlers"
)

// Will import http endpoints for the main router from internal packages.
func ConfigureRoutes(router *http.ServeMux, h *handlers.Handler) {

	// Base routes
	handlers.BaseRoutes(router, h)
}
