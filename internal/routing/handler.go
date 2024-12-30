package routing

import (
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/LoganDarrinLee/market-ctf/internal/common"
)

type Handler struct {
	Pool   *pgxpool.Pool
	Logger *common.BasicLogger
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

func (h *Handler) authHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Authenticated page."))
	})
}
