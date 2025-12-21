package services

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type AboutService struct {
	Repo *repository.Repository
	Log  *zap.Logger
}

func NewAboutService(repo *repository.Repository, log *zap.Logger) *AboutService {
	return &AboutService{
		Repo: repo,
		Log:  log,
	}
}
