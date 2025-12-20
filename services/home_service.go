package services

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type HomeService struct {
	Repo   *repository.Repository
	Logger *zap.Logger
}

func NewHomeService(repo *repository.Repository, log *zap.Logger) *HomeService {
	return &HomeService{
		Repo:   repo,
		Logger: log,
	}
}
