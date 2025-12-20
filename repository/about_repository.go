package repository

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"go.uber.org/zap"
)

type AboutRepository struct {
	DB  db.PgxIface
	Log *zap.Logger
}

func NewAboutRepository(db db.PgxIface, log *zap.Logger) *AboutRepository {
	return &AboutRepository{
		DB:  db,
		Log: log,
	}
}
