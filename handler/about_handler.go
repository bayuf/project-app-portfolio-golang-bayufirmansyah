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
	profile, err := h.Service.GetService.GetProfile()
	if err != nil {
		h.Logger.Error("cant get profil data", zap.Error(err))
	}

	skills, err := h.Service.GetService.GetAllSkills()
	if err != nil {
		h.Logger.Error("cant get skills data", zap.Error(err))
	}

	feedbacks, err := h.Service.GetService.GetAllFeedbacks()
	if err != nil {
		h.Logger.Error("cant get feedback data", zap.Error(err))
	}

	data := Data{
		Title:     "About me",
		Profile:   profile,
		Skills:    skills,
		Feedbacks: feedbacks,
		Nav:       BuildNav("About"),
	}

	if err := h.Template.ExecuteTemplate(w, "about-page", data); err != nil {
		h.Logger.Error("failed to execute about page template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
