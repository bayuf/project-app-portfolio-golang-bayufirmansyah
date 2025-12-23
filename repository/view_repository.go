package repository

import (
	"context"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type ViewRepository struct {
	DB     db.PgxIface
	Logger *zap.Logger
}

func NewViewRepository(db db.PgxIface, log *zap.Logger) *ViewRepository {
	return &ViewRepository{
		DB:     db,
		Logger: log,
	}
}

func (r *ViewRepository) GetProfile() (*model.Profile, error) {
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

func (r *ViewRepository) GetAddress() (*model.Address, error) {
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

func (r *ViewRepository) GetAllSkills() (*[]model.Skill, error) {
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

func (r *ViewRepository) GetAllOffers() (*[]model.Offers, error) {
	query := `SELECT id, user_id, title, description, icon_url, created_at, updated_at
			  FROM offers
			  WHERE deleted_at IS NULL;`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		r.Logger.Error("query getAllOffers failed", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var offers []model.Offers
	for rows.Next() {
		offer := model.Offers{}
		rows.Scan(&offer.Model.ID, &offer.UserId, &offer.Title, &offer.Description, &offer.IconURL,
			&offer.Model.Created_At, &offer.Model.Updated_At)

		offers = append(offers, offer)
	}

	return &offers, nil

}

func (r *ViewRepository) GetAllProject() (*[]model.Project, error) {
	query := `SELECT id, user_id, title, description, icon_url, link_url, created_at, updated_at
			  FROM projects
			  WHERE deleted_at IS NULL;`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		r.Logger.Error("query getAllProjects failed", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		project := model.Project{}
		rows.Scan(&project.Model.ID, &project.UserId, &project.Title, &project.Description, &project.IconURL,
			&project.LinkURL, &project.Model.Created_At, &project.Model.Updated_At)

		projects = append(projects, project)
	}

	return &projects, nil

}

func (r *ViewRepository) GetAllFeedbacks() (*[]model.Feedback, error) {
	query := `SELECT id, name, message, profile_img, created_at
			  FROM feedbacks
			  WHERE deleted_at IS NULL;`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		r.Logger.Error("query getAllfeedback failed", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var feedbacks []model.Feedback
	for rows.Next() {
		feedback := model.Feedback{}
		rows.Scan(&feedback.Model.ID, &feedback.Name, &feedback.Message, &feedback.ProfileURL, &feedback.Model.Created_At)

		feedbacks = append(feedbacks, feedback)
	}

	return &feedbacks, nil

}
