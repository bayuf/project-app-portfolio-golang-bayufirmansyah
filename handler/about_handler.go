package handler

import (
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type AboutHandler struct {
	Service  *services.Service
	Template *template.Template
	Logger   *zap.Logger
}

func NewAboutHandler(service *services.Service, template *template.Template, log *zap.Logger) *AboutHandler {
	return &AboutHandler{
		Service:  service,
		Template: template,
		Logger:   log,
	}
}

func (h *AboutHandler) AboutpageView(w http.ResponseWriter, r *http.Request) {
	if err := h.Template.ExecuteTemplate(w, "about", nil); err != nil {
		h.Logger.Error("failed to execute about page template", zap.Error(err))
	}
}
