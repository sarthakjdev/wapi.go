package utils

import (
	"github.com/go-playground/validator/v10"
)

// return a validator instance
func GetValidator() *validator.Validate {
	return validator.New()
}
