package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
)

type Handler struct {
	Service  *services.Service
	Template *template.Template
}

func NewHandler(service *services.Service, template *template.Template) *Handler {
	return &Handler{
		Service:  service,
		Template: template,
	}
}

func (h *Handler) HomepageView(w http.ResponseWriter, r *http.Request) {
	if err := h.Template.ExecuteTemplate(w, "/", nil); err != nil {
		log.Println(err)
	}
}
