package repository

import (
	"context"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"go.uber.org/zap"
)

type InputRepository struct {
	DB     db.PgxIface
	Logger *zap.Logger
}

func NewInputRepository(db db.PgxIface, log *zap.Logger) *InputRepository {
	return &InputRepository{
		DB:     db,
		Logger: log,
	}
}

func (r *InputRepository) SendMessage(data *model.Message) error {
	query := `INSERT INTO messages (name, email, subject, message)
	VALUES ($1, $2, $3, $4);`

	_, err := r.DB.Exec(context.Background(), query, data.Name, data.Email, data.Subject, data.Message)
	if err != nil {
		r.Logger.Error("cant send message", zap.Error(err))
		return err
	}

	return nil
}
