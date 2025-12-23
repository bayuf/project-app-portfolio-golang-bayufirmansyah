package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateMessage(dto any) error {
	return validate.Struct(dto)
}
