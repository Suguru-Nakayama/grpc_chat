package validation

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func validPassword(fl validator.FieldLevel) bool {
	var low, num bool

	password := fl.Field().String()

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			low = true
		case unicode.IsNumber(char):
			num = true
		default:
			return false
		}
	}

	if !low || !num || len(password) < 8 || len(password) > 32 {
		return false
	}
	return true
}
