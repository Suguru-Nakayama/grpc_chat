package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type LogInValidator struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func NewLogInValidator(email, password string) *LogInValidator {
	return &LogInValidator{
		Email:    email,
		Password: password,
	}
}

func (luv *LogInValidator) Validate() map[string]string {
	v := validator.New()
	errors := make(map[string]string, 0)
	if err := v.Struct(luv); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var msg string
			field := string(err.Field())
			switch field {
			case "Email":
				switch err.ActualTag() {
				case "required":
					msg = fmt.Sprintf("メールアドレスを入力してください")
				}
			case "Password":
				switch err.ActualTag() {
				case "required":
					msg = fmt.Sprintf("パスワードを入力してください")
				}
			}
			errors[field] = msg
		}
	}
	return errors
}
