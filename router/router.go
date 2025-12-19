package router

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/handler"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"github.com/go-chi/chi/v5"
)

func NewRouter(svc *services.Service, handl *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	// root
	r.Get("/", handl.HomepageView)

	return r
}
