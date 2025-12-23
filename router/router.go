package router

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/handler"
	customMiddleware "github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/middleware"
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

	// Admin

	r.Get("/login", handl.AuthHandler.LoginView)
	r.Post("/login", handl.AuthHandler.Login)

	// root
	r.Get("/", handl.HomeHandler.HomepageView)

	// Download
	r.Route("/download", func(r chi.Router) {
		r.Get("/cv", handl.DownloadHandler.DownloadCV)
		r.Get("/cv-detail", handl.DownloadHandler.SeeCV)

	})

	// pages
	// About
	r.Get("/about", handl.AboutHandler.AboutpageView)

	// Services
	r.Get("/services", handl.ServicesHandler.ServicesPageView)

	// Portofolio
	r.Get("/portofolio", handl.PortofolioHandler.PortofolioPageView)

	// Contact
	r.Route("/contact", func(r chi.Router) {
		r.Get("/", handl.ContactHandler.ContactPageView)
		r.Post("/send", handl.ContactHandler.SendMessageHandler)
	})
	// admin pages protected
	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.AuthMiddleware)
		r.Route("/admin", func(r chi.Router) {
			r.Get("/dashboard", handl.DashboardHandler.DashboardView)
		})

	})

	return r
}
