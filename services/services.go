package services

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type Service struct {
	Repo   *repository.Repository
	Logger *zap.Logger

	// cache
	Profil *model.Profile
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		Repo:   repo,
		Logger: log,
	}
}

func (s *Service) GetProfile() (*model.Profile, error) {
	if s.Profil != nil {
		return s.Profil, nil
	}

	// get from db
	return s.Repo.GetProfile()
}
