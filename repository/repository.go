package repository

import (
	"context"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type Repository struct {
	DB     db.PgxIface
	Logger *zap.Logger
}

func NewRepository(db db.PgxIface, log *zap.Logger) *Repository {
	return &Repository{
		DB:     db,
		Logger: log,
	}
}

func (r *Repository) GetProfile() (*model.Profile, error) {
	query := `SELECT id, first_name, last_name, headline, about, email, phone, created_at, updated_at
			 FROM profiles
			 WHERE deleted_at IS NULL`

	row := r.DB.QueryRow(context.Background(), query)

	var profile model.Profile
	if err := row.Scan(
		&profile.Model.ID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Headline,
		&profile.About,
		&profile.Email,
		&profile.Phone,
		&profile.Model.Created_At,
		&profile.Model.Updated_At,
	); err != nil {
		if err == pgx.ErrNoRows {
			r.Logger.Info("id not found in profile", zap.Error(err))
		} else {
			r.Logger.Error("error in repository GetProfileById", zap.Error(err))
		}
		return nil, err
	}

	return &profile, nil
}
