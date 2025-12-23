package handler

import (
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type DashboardHandler struct {
	Service  *services.Service
	Template *template.Template
	Logger   *zap.Logger
}

func NewDashboardHandler(service *services.Service, template *template.Template, log *zap.Logger) *DashboardHandler {
	return &DashboardHandler{
		Service:  service,
		Template: template,
		Logger:   log,
	}
}

func (h *DashboardHandler) DashboardView(w http.ResponseWriter, r *http.Request) {
	if err := h.Template.ExecuteTemplate(w, "dashboard", nil); err != nil {
		h.Logger.Error("failed to execute dashboard template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
