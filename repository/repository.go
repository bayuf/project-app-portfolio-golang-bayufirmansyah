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
	query := `SELECT id, first_name, last_name, headline, about, email, phone, created_at, updated_at, experience
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
		&profile.Experience,
	); err != nil {
		if err == pgx.ErrNoRows {
			r.Logger.Info("profile is null", zap.Error(err))
		} else {
			r.Logger.Error("error in repository GetProfile", zap.Error(err))
		}
		return nil, err
	}

	return &profile, nil
}

func (r *Repository) GetAddress() (*model.Address, error) {
	query := `SELECT id, user_id, address, city, country, province, postal_code,
		created_at, updated_at
		FROM address
		WHERE deleted_at IS NULL;`

	row := r.DB.QueryRow(context.Background(), query)
	var address model.Address
	if err := row.Scan(
		&address.Model.ID,
		&address.UserId,
		&address.Detail,
		&address.City,
		&address.Country,
		&address.Province,
		&address.PostalCode,
		&address.Model.Created_At,
		&address.Model.Updated_At,
	); err != nil {
		if err == pgx.ErrNoRows {
			r.Logger.Info("address is null", zap.Error(err))
		} else {
			r.Logger.Error("error in repository GetAddress", zap.Error(err))
		}
		return nil, err
	}

	return &address, nil

}

func (r *Repository) GetAllSkills() (*[]model.Skill, error) {
	query := `SELECT id, user_id, name, level, img_url, created_at, updated_at
		FROM skills
		WHERE deleted_at IS NULL;`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		r.Logger.Error("query getAllSkill failed", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var skills []model.Skill
	for rows.Next() {
		skill := model.Skill{}
		rows.Scan(&skill.Model.ID, &skill.UserId, &skill.Name, &skill.Level, &skill.ImgURL,
			&skill.Model.Created_At, &skill.Model.Updated_At)

		skills = append(skills, skill)
	}

	return &skills, nil
}
