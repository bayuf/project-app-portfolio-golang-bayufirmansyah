package services

import (
	"errors"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/dto"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/utils"
	"go.uber.org/zap"
)

type AuthService struct {
	Repo   *repository.AuthRepository
	Logger *zap.Logger
}

func NewAuthService(repo *repository.AuthRepository, log *zap.Logger) *AuthService {
	return &AuthService{
		Repo:   repo,
		Logger: log,
	}
}

func (s *AuthService) Login(data *dto.UserDTO) (*model.User, error) {

	user, err := s.Repo.GetUserDataByEmail(&model.User{
		Email: data.Email,
	})
	if err != nil {
		s.Logger.Error("service cant get user data", zap.Error(err))
		return &model.User{}, errors.New("credentials error")
	}

	if !utils.CheckPassword(user.Password, data.Password) {
		s.Logger.Info("incorect password")
		return nil, errors.New("credentials error")
	}

	return user, nil
}
