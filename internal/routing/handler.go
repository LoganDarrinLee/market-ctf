package routing

import (
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/LoganDarrinLee/market-ctf/internal/common"
	"github.com/LoganDarrinLee/market-ctf/internal/db"
)

type Handler struct {
	Queries *db.Queries
	Pool    *pgxpool.Pool
	Logger  *common.BasicLogger
}

func NewHandler(l *common.BasicLogger) *Handler {
	return &Handler{Logger: l}
}

// Index page. Display general information
func (h *Handler) indexHandler(w http.ResponseWriter, r *http.Request) {
	// Access level,

	// Extract request context
	rc := common.NewRequestContext(r.Context())

	// Log request info.
	// h.Logger.WriteInfo(rc, "GET /")

	// Page data
	data := make(map[string]interface{})
	data["title"] = "Index Title"
	data["msg"] = []string{"test1", "second", "third", "fourth"}

	// html template
	tmpl, err := template.ParseFiles(
		"web/templates/base.html", "web/templates/index.html")
	if err != nil {
		h.Logger.WriteError(rc, "Error parsing template files", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Execute template with page data
	tmpl.Execute(w, data)
}

func (h *Handler) aboutHandler(w http.ResponseWriter, r *http.Request) {
	// Extract request context
	rc := common.NewRequestContext(r.Context())

	// Log request info.
	h.Logger.WriteInfo(rc, "GET /")

	data := make(map[string]interface{})
	data["title"] = "About title"

	// Create page template.
	tmpl, err := template.ParseFiles(
		"web/templates/base.html", "web/templates/about.html")
	if err != nil {
		h.Logger.WriteError(rc, "Error parsing template files", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl.Execute(w, data)
}

func (h *Handler) authHandler(w http.ResponseWriter, r *http.Request) {
	// Request context
	rc := common.NewRequestContext(r.Context())

	// Check authentication.
	authenticated, userSessionRow, err := checkAuth(h.Queries, rc, "auth")
	if err != nil {
		h.Logger.WriteError(rc, "Error checking authentication.", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// If the user is authenticated
	if authenticated {
		h.Logger.WriteInfo(rc, userSessionRow.PublicUsername)
		w.Write([]byte("Authenticated"))
	} else {
		// The user was not authenticated.
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	}
}

// The login page handler will process user login attempts, or if they already
// contain a valid session token, redirect to homepage.
func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	// New request context object.
	rc := common.NewRequestContext(r.Context())

	// Check for user authentication before processing a login.
	if isAuthenticated() {
		// User session token exists and is valid.
		w.Write([]byte("Welcome to the market."))
	} else {
		// Proceed with rendering login information.
		w.Write([]byte("Please proceed with login."))
	}

}
