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
	Profil  *model.Profile
	Address *model.Address
	Skills  *[]model.Skill
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		Repo:   repo,
		Logger: log,
	}
}

func (s *Service) GetProfile() (*model.Profile, error) {
	// cache check
	if s.Profil != nil {
		return s.Profil, nil
	}

	var err error
	s.Profil, err = s.Repo.GetProfile()
	if err != nil {
		s.Logger.Error("service error cant get profile", zap.Error(err))
		return nil, err
	}

	return s.Profil, nil
}

func (s *Service) GetAddress() (*model.Address, error) {
	if s.Address != nil {
		return s.Address, nil
	}

	var err error
	s.Address, err = s.Repo.GetAddress()
	if err != nil {
		s.Logger.Error("service error cant get address", zap.Error(err))
		return nil, err
	}

	return s.Address, nil
}

func (s *Service) GetAllSkills() (*[]model.Skill, error) {
	if s.Skills != nil {
		return s.Skills, nil
	}

	var err error
	s.Skills, err = s.Repo.GetAllSkills()
	if err != nil {
		s.Logger.Error("service error cant get skills", zap.Error(err))
		return nil, err
	}

	return s.Skills, nil
}
