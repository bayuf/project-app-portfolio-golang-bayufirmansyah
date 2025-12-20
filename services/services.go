package services

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type Service struct {
	HomeService *HomeService
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		HomeService: NewHomeService(repo, log),
	}
}
