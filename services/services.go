package services

import (
	"errors"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/dto"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"go.uber.org/zap"
)

type Service struct {
	Repo   *repository.Repository
	Logger *zap.Logger

	// cache
	Profil    *model.Profile
	Address   *model.Address
	Skills    *[]model.Skill
	Offers    *[]model.Offers
	Projects  *[]model.Project
	Feedbacks *[]model.Feedback
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		Repo:   repo,
		Logger: log,
	}
}

func (s *Service) Login(data *dto.UserDTO) (*model.User, error) {
	user, err := s.Repo.AuthRepository.GetUserDataByEmail(&model.User{
		Email: data.Email,
	})
	if err != nil {
		s.Logger.Error("service cant get user data", zap.Error(err))
	}

	if user.Password != data.Password {
		s.Logger.Info("incorect password")
		return nil, errors.New("incorect password")
	}

	return user, nil
}

func (s *Service) GetProfile() (*model.Profile, error) {
	// cache check
	if s.Profil != nil {
		return s.Profil, nil
	}

	var err error
	s.Profil, err = s.Repo.ViewRepository.GetProfile()
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
	s.Address, err = s.Repo.ViewRepository.GetAddress()
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
	s.Skills, err = s.Repo.ViewRepository.GetAllSkills()
	if err != nil {
		s.Logger.Error("service error cant get skills", zap.Error(err))
		return nil, err
	}

	return s.Skills, nil
}

func (s *Service) GetAllOffers() (*[]model.Offers, error) {
	if s.Offers != nil {
		return s.Offers, nil
	}

	var err error
	s.Offers, err = s.Repo.ViewRepository.GetAllOffers()
	if err != nil {
		s.Logger.Error("service error cant get offers", zap.Error(err))
		return nil, err
	}

	return s.Offers, nil
}

func (s *Service) GetAllProjects() (*[]model.Project, error) {
	if s.Projects != nil {
		return s.Projects, nil
	}

	var err error
	s.Projects, err = s.Repo.ViewRepository.GetAllProject()
	if err != nil {
		s.Logger.Error("service error cant get projects", zap.Error(err))
		return nil, err
	}

	return s.Projects, nil
}

func (s *Service) GetAllFeedbacks() (*[]model.Feedback, error) {
	if s.Feedbacks != nil {
		return s.Feedbacks, nil
	}

	var err error
	s.Feedbacks, err = s.Repo.ViewRepository.GetAllFeedbacks()
	if err != nil {
		s.Logger.Error("service error cant get feedbacks", zap.Error(err))
		return nil, err
	}

	return s.Feedbacks, nil
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
