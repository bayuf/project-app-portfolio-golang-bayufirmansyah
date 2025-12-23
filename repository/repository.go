package repository

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"go.uber.org/zap"
)

type Repository struct {
	ViewRepository  ViewRepository
	InputRepository InputRepository
	AuthRepository  AuthRepository
}

func NewRepository(db db.PgxIface, log *zap.Logger) *Repository {
	return &Repository{
		ViewRepository:  *NewViewRepository(db, log),
		InputRepository: *NewInputRepository(db, log),
		AuthRepository:  *NewAuthRepository(db, log),
	}
}
