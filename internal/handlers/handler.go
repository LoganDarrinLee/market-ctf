package handlers

import (
	"html/template"
	"log"
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

// Pick my brain page. Convert vault to knowledge
func (h *Handler) brainPage() {}

// Index page. Display general information
func (h *Handler) indexHandler(w http.ResponseWriter, r *http.Request) {
	// Extract request context
	rc := common.NewRequestContext(r.Context())

	// Log request info.
	h.Logger.WriteInfo(rc, "GET /")

	data := make(map[string]interface{})
	data["title"] = "Index Title"
	data["msg"] = []string{"test1", "second", "third", "fourth"}

	tmpl, err := template.ParseFiles(
		"web/templates/base.html", "web/templates/index.html")
	if err != nil {
		h.Logger.WriteError(rc, "Error parsing template files", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println(r.Context().Value(common.RequestIDKey))
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
