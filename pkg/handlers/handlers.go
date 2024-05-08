package handlers

import (
	"net/http"

	"github.com/shaynemeyer/go-bnb/pkg/config"
	"github.com/shaynemeyer/go-bnb/pkg/models"
	"github.com/shaynemeyer/go-bnb/pkg/render"
)

// Repo - Repository used by the handlers
var Repo *Repository

// Repository - Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo - creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers - sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
