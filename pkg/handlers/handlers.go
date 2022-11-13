package handlers

import (
	"fmt"
	"github.com/dorukbulut/bookings/pkg/config"
	"github.com/dorukbulut/bookings/pkg/models"
	"github.com/dorukbulut/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteID := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteID)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, test"

	remoteIp := m.App.Session.Get(r.Context(), "remote_ip")

	stringMap["remote_ip"] = fmt.Sprintf("%v", remoteIp)
	// render the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
