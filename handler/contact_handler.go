package handler

import (
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type ContactHandler struct {
	Service  *services.Service
	Template *template.Template
	Logger   *zap.Logger
}

func NewContactHandler(service *services.Service, template *template.Template, log *zap.Logger) *ContactHandler {
	return &ContactHandler{
		Service:  service,
		Template: template,
		Logger:   log,
	}
}

func (h *ContactHandler) ContactPageView(w http.ResponseWriter, r *http.Request) {
	if err := h.Template.ExecuteTemplate(w, "contact-page", nil); err != nil {
		h.Logger.Error("failed to execute contact page template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
