package repository

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"go.uber.org/zap"
)

type Repository struct {
	DB        db.PgxIface
	HomeRepo  *HomeRepository
	AboutRepo *AboutRepository
}

func NewRepository(db db.PgxIface, log *zap.Logger) *Repository {
	return &Repository{
		HomeRepo:  NewHomeRepository(db, log),
		AboutRepo: NewAboutRepository(db, log),
	}
}
