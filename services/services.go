package services

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/dto"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type Service struct {
	GetService  *GetService
	AuthService *AuthService
	Repo        *repository.Repository
	Logger      *zap.Logger
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		GetService:  NewGetService(&repo.ViewRepository, log),
		AuthService: NewAuthService(&repo.AuthRepository, log),
	}
}

func (s *Service) SendMessage(data *dto.MessageDTO) error {
	if err := s.Repo.InputRepository.SendMessage(&model.Message{
		Name:    data.Name,
		Email:   data.Email,
		Subject: data.Subject,
		Message: data.Message,
	}); err != nil {
		s.Logger.Error("service send message failed", zap.Error(err))
		return err
	}
	return nil
}
