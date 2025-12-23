package repository

import (
	"context"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type AuthRepository struct {
	DB     db.PgxIface
	Logger *zap.Logger
}

func NewAuthRepository(db db.PgxIface, log *zap.Logger) *AuthRepository {
	return &AuthRepository{
		DB:     db,
		Logger: log,
	}
}

func (r *AuthRepository) GetUserDataByEmail(email *model.User) (*model.User, error) {
	query := `SELECT id, email, password_hash, last_login_at, created_at, updated_at
			  FROM auth_users
			  WHERE email=$1`

	row := r.DB.QueryRow(context.Background(), query, email.Email)

	var user model.User
	if err := row.Scan(&user.Model.ID, &user.Email, &user.Password, &user.LastLogin,
		&user.Model.Created_At, &user.Model.Updated_At); err != nil {
		if err == pgx.ErrNoRows {
			r.Logger.Info("email not found", zap.Error(err))
		} else {
			r.Logger.Error("error in repository GetUser", zap.Error(err))
		}
		return nil, err

	}
	return &user, nil
}
