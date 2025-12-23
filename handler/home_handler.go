package handler

import (
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type HomeHandler struct {
	Service  *services.Service
	Template *template.Template
	Logger   *zap.Logger
}

func NewHomeHandler(service *services.Service, template *template.Template, log *zap.Logger) *HomeHandler {
	return &HomeHandler{
		Service:  service,
		Template: template,
		Logger:   log,
	}
}

func (h *HomeHandler) HomepageView(w http.ResponseWriter, r *http.Request) {
	profile, err := h.Service.GetProfile()
	if err != nil {
		h.Logger.Error("cant get profil data", zap.Error(err))
	}
	skills, err := h.Service.GetAllSkills()
	if err != nil {
		h.Logger.Error("cant get skills data", zap.Error(err))
	}

	data := Data{
		Title:   "Homepage",
		Profile: profile,
		Skills:  skills,
		Nav:     BuildNav("Home"),
	}

	if err := h.Template.ExecuteTemplate(w, "/", data); err != nil {
		h.Logger.Error("failed to execute homepage template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
