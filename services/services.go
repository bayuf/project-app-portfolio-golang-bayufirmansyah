package services

import "github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}
