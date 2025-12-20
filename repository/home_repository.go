package repository

import (
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"go.uber.org/zap"
)

type HomeRepository struct {
	DB  db.PgxIface
	Log *zap.Logger
}

func NewHomeRepository(db db.PgxIface, log *zap.Logger) *HomeRepository {
	return &HomeRepository{
		DB:  db,
		Log: log,
	}
}
