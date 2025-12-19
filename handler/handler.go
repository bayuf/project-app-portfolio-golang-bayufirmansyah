package handler

import (
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type Handler struct {
	HomeHandler  *HomeHandler
	AboutHandler *AboutHandler
}

func NewHandler(service *services.Service, template *template.Template, logger *zap.Logger) *Handler {
	return &Handler{
		HomeHandler:  NewHomeHandler(service, template, logger),
		AboutHandler: NewAboutHandler(service, template, logger),
	}
}
