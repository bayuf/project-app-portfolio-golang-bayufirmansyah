package handler

import (
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"go.uber.org/zap"
)

type Handler struct {
	HomeHandler       *HomeHandler
	AboutHandler      *AboutHandler
	ServicesHandler   *ServicesHandler
	PortofolioHandler *PortofolioHandler
	ContactHandler    *ContactHandler
	DownloadHandler   *DownloadHandler
	AuthHandler       *AuthHandler
	DashboardHandler  *DashboardHandler
}

type Data struct {
	Title          string
	Success, Error bool
	Profile        *model.Profile
	Address        *model.Address
	Skills         *[]model.Skill
	Offers         *[]model.Offers
	Projects       *[]model.Project
	Feedbacks      *[]model.Feedback
	Nav            []model.NavItem
}

func NewHandler(service *services.Service, template *template.Template, logger *zap.Logger) *Handler {
	return &Handler{
		HomeHandler:       NewHomeHandler(service, template, logger),
		AboutHandler:      NewAboutHandler(service, template, logger),
		ServicesHandler:   NewServicesHandler(service, template, logger),
		PortofolioHandler: NewPortofolioHandler(service, template, logger),
		ContactHandler:    NewContactHandler(service, template, logger),
		DashboardHandler:  NewDashboardHandler(service, template, logger),

		DownloadHandler: NewDownloadHandler(logger),

		AuthHandler: NewAuthHandler(service, template, logger),
	}
}
