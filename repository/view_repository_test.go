package repository_test

import (
	"testing"
	"time"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetProfile_Success(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockDB.Close()

	logger := zap.NewNop()
	repo := repository.NewViewRepository(mockDB, logger)

	rows := pgxmock.NewRows([]string{
		"id", "first_name", "last_name", "headline", "about",
		"email", "phone", "created_at", "updated_at", "experience",
	}).AddRow(
		1,
		"Bayu",
		"Firmansyah",
		"Backend Engineer",
		"About me",
		"bayu@email.com",
		"08123456789",
		time.Now(),
		time.Now(),
		5,
	)

	mockDB.ExpectQuery("FROM profiles").
		WillReturnRows(rows)

	profile, err := repo.GetProfile()

	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, "Bayu", profile.FirstName)
	assert.Equal(t, 5, profile.Experience)
}

func TestGetProfile_NoRows(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockDB.Close()

	logger := zap.NewNop()
	repo := repository.NewViewRepository(mockDB, logger)

	// rows kosong (tidak ada AddRow)
	rows := pgxmock.NewRows([]string{
		"id", "first_name", "last_name", "headline", "about",
		"email", "phone", "created_at", "updated_at", "experience",
	})

	mockDB.ExpectQuery("FROM profiles").
		WillReturnRows(rows)

	profile, err := repo.GetProfile()

	assert.Error(t, err)
	assert.Nil(t, profile)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}

func TestGetAllSkills_Success(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockDB.Close()

	logger := zap.NewNop()
	repo := repository.NewViewRepository(mockDB, logger)

	rows := pgxmock.NewRows([]string{
		"id", "user_id", "name", "level", "img_url", "created_at", "updated_at",
	}).AddRow(
		1, 1, "Golang", "advanced", "icon.png", time.Now(), time.Now(),
	).AddRow(
		2, 1, "PostgreSQL", "advanced", "icon2.png", time.Now(), time.Now(),
	)

	mockDB.ExpectQuery("FROM skills").
		WillReturnRows(rows)

	skills, err := repo.GetAllSkills()

	assert.NoError(t, err)
	assert.Len(t, *skills, 2)
	assert.Equal(t, "Golang", (*skills)[0].Name)
}

func TestGetAllOffers_Success(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockDB.Close()

	logger := zap.NewNop()
	repo := repository.NewViewRepository(mockDB, logger)

	rows := pgxmock.NewRows([]string{
		"id", "user_id", "title", "description", "icon_url", "created_at", "updated_at",
	}).AddRow(
		1, 1, "Backend", "Backend Golang", "icon.png", time.Now(), time.Now(),
	)

	mockDB.ExpectQuery("FROM offers").
		WillReturnRows(rows)

	offers, err := repo.GetAllOffers()

	assert.NoError(t, err)
	assert.Len(t, *offers, 1)
	assert.Equal(t, "Backend", (*offers)[0].Title)
}

func TestGetAllFeedbacks_Success(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mockDB.Close()

	logger := zap.NewNop()
	repo := repository.NewViewRepository(mockDB, logger)

	rows := pgxmock.NewRows([]string{
		"id", "name", "message", "profile_img", "created_at",
	}).AddRow(
		1, "Andi", "Bagus sekali", "user.png", time.Now(),
	)

	mockDB.ExpectQuery("FROM feedbacks").
		WillReturnRows(rows)

	feedbacks, err := repo.GetAllFeedbacks()

	assert.NoError(t, err)
	assert.Len(t, *feedbacks, 1)
	assert.Equal(t, "Andi", (*feedbacks)[0].Name)
}
