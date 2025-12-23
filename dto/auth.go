package dto

type UserDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
