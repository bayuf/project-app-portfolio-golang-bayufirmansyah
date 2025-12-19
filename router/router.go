package router

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/handler"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewRouter(svc *services.Service, handl *handler.Handler, log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	// chi middleware
	// Logger
	r.Use(middleware.Logger)
	// Panic Recover
	r.Use(middleware.Recoverer)

	// root
	r.Get("/", handl.HomeHandler.HomepageView)

	// pages
	// About
	r.Get("/about", handl.AboutHandler.AboutpageView)

	return r
}
