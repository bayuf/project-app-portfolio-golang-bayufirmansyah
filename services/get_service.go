package services

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type GetService struct {
	Repo   *repository.ViewRepository
	Logger *zap.Logger
	Cache
}

type Cache struct {
	// cache
	Profil    *model.Profile
	Address   *model.Address
	Skills    *[]model.Skill
	Offers    *[]model.Offers
	Projects  *[]model.Project
	Feedbacks *[]model.Feedback
}

func NewGetService(repo *repository.ViewRepository, log *zap.Logger) *GetService {
	return &GetService{
		Repo:   repo,
		Logger: log,
	}
}

func (s *GetService) GetProfile() (*model.Profile, error) {
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

func (s *GetService) GetAddress() (*model.Address, error) {
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

func (s *GetService) GetAllSkills() (*[]model.Skill, error) {
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

func (s *GetService) GetAllOffers() (*[]model.Offers, error) {
	if s.Offers != nil {
		return s.Offers, nil
	}

	var err error
	s.Offers, err = s.Repo.GetAllOffers()
	if err != nil {
		s.Logger.Error("service error cant get offers", zap.Error(err))
		return nil, err
	}

	return s.Offers, nil
}

func (s *GetService) GetAllProjects() (*[]model.Project, error) {
	if s.Projects != nil {
		return s.Projects, nil
	}

	var err error
	s.Projects, err = s.Repo.GetAllProject()
	if err != nil {
		s.Logger.Error("service error cant get projects", zap.Error(err))
		return nil, err
	}

	return s.Projects, nil
}

func (s *GetService) GetAllFeedbacks() (*[]model.Feedback, error) {
	if s.Feedbacks != nil {
		return s.Feedbacks, nil
	}

	var err error
	s.Feedbacks, err = s.Repo.GetAllFeedbacks()
	if err != nil {
		s.Logger.Error("service error cant get feedbacks", zap.Error(err))
		return nil, err
	}

	return s.Feedbacks, nil
}
