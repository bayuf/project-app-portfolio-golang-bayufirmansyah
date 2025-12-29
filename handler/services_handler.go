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
	profile, err := h.Service.GetService.GetProfile()
	if err != nil {
		h.Logger.Error("cant get profil data", zap.Error(err))
	}

	offers, err := h.Service.GetService.GetAllOffers()
	if err != nil {
		h.Logger.Error("cant get offers data", zap.Error(err))
	}

	feedbacks, err := h.Service.GetService.GetAllFeedbacks()
	if err != nil {
		h.Logger.Error("cant get feedback data", zap.Error(err))
	}

	data := Data{
		Title:     "Services",
		Profile:   profile,
		Offers:    offers,
		Feedbacks: feedbacks,
		Nav:       BuildNav("Services"),
	}
	if err := h.Template.ExecuteTemplate(w, "services-page", data); err != nil {
		h.Logger.Error("failed to execute services page template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
