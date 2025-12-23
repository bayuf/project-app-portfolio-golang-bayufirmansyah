package dto

type MessageDTO struct {
	Name    string `validate:"required,min=2,max=75"`
	Email   string `validate:"required,email,max=50"`
	Subject string `validate:"required,min=3,max=150"`
	Message string `validate:"required,min=5,max=1000"`
}
