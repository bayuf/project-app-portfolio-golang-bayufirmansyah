package handler

import (
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type PortofolioHandler struct {
	Service  *services.Service
	Template *template.Template
	Logger   *zap.Logger
}

func NewPortofolioHandler(service *services.Service, template *template.Template, log *zap.Logger) *PortofolioHandler {
	return &PortofolioHandler{
		Service:  service,
		Template: template,
		Logger:   log,
	}
}

func (h *PortofolioHandler) PortofolioPageView(w http.ResponseWriter, r *http.Request) {
	if err := h.Template.ExecuteTemplate(w, "portofolio-page", nil); err != nil {
		h.Logger.Error("failed to execute portofolio page template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
