package handler

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/dto"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/utils"
	"go.uber.org/zap"
)

type ContactHandler struct {
	Service  *services.Service
	Template *template.Template
	Logger   *zap.Logger
}

func NewContactHandler(service *services.Service, template *template.Template, log *zap.Logger) *ContactHandler {
	return &ContactHandler{
		Service:  service,
		Template: template,
		Logger:   log,
	}
}

func (h *ContactHandler) ContactPageView(w http.ResponseWriter, r *http.Request) {
	success := r.URL.Query().Get("success") == "1"
	error := r.URL.Query().Get("error") == "1"

	profile, err := h.Service.GetProfile()
	if err != nil {
		h.Logger.Error("cant get profil data", zap.Error(err))
	}

	address, err := h.Service.GetAddress()
	if err != nil {
		h.Logger.Error("cant get address data", zap.Error(err))
	}

	data := Data{
		Title:   "Contact",
		Success: success,
		Error:   error,
		Profile: profile,
		Address: address,
		Nav:     BuildNav("Contact"),
	}

	if err := h.Template.ExecuteTemplate(w, "contact-page", data); err != nil {
		h.Logger.Error("failed to execute contact page template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *ContactHandler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.Logger.Error("try to access send message with another method", zap.String("method:", r.Method))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// get data from form
	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))
	subject := strings.TrimSpace(r.FormValue("subject"))
	message := strings.TrimSpace(r.FormValue("message"))

	// validate
	messageData := dto.MessageDTO{
		Name:    name,
		Email:   email,
		Subject: subject,
		Message: message,
	}
	if err := utils.ValidateMessage(&messageData); err != nil {
		h.Logger.Error("input for message invalid", zap.Error(err))
		http.Redirect(w, r, "/contact?error=1", http.StatusSeeOther)
		return
	}

	// call service
	if err := h.Service.SendMessage(&messageData); err != nil {
		h.Logger.Error("service error", zap.Error(err))
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contact?success=1", http.StatusSeeOther)
}
