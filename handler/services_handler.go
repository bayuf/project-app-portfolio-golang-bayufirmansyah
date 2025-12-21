package handler

import (
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type ServicesHandler struct {
	Service  *services.Service
	Template *template.Template
	Logger   *zap.Logger
}

func NewServicesHandler(service *services.Service, template *template.Template, log *zap.Logger) *ServicesHandler {
	return &ServicesHandler{
		Service:  service,
		Template: template,
		Logger:   log,
	}
}

func (h *ServicesHandler) ServicesPageView(w http.ResponseWriter, r *http.Request) {
	if err := h.Template.ExecuteTemplate(w, "services-page", nil); err != nil {
		h.Logger.Error("failed to execute services page template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
