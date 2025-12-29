package handler

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/dto"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/utils"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service  *services.AuthService
	Template *template.Template
	Logger   *zap.Logger
}

func NewAuthHandler(service *services.Service, template *template.Template, log *zap.Logger) *AuthHandler {
	return &AuthHandler{
		Service:  service.AuthService,
		Template: template,
		Logger:   log,
	}
}

func (h *AuthHandler) LoginView(w http.ResponseWriter, r *http.Request) {
	if err := h.Template.ExecuteTemplate(w, "login", nil); err != nil {
		h.Logger.Error("failed to execute login page template", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
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
	email := strings.TrimSpace(r.FormValue("email"))
	password := strings.TrimSpace(r.FormValue("password"))

	userData := dto.UserDTO{
		Email:    email,
		Password: password,
	}

	if err := utils.ValidateMessage(userData); err != nil {
		h.Logger.Error("input for message invalid", zap.Error(err))
		http.Error(w, "invalid input data", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Login(&userData)
	if err != nil {
		http.Error(w, "login failed", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    strconv.Itoa(user.Model.ID),
		Path:     "/",
		HttpOnly: true,
	})
	http.Redirect(w, r, "admin/dashboard", http.StatusSeeOther)

}
