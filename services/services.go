package services

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type Service struct {
	HomeService  *HomeService
	AboutService *AboutService
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		HomeService:  NewHomeService(repo, log),
		AboutService: NewAboutService(repo, log),
	}
}
