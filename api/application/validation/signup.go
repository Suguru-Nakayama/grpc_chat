package validation

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	LastName  string `validate:"required,min=1,max=25"`
	FirstName string `validate:"required,min=1,max=25"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,password"`
}

func NewUserValidator(lastName, firstName, email, password string) *UserValidator {
	return &UserValidator{
		LastName:  lastName,
		FirstName: firstName,
		Email:     email,
		Password:  password,
	}
}

func (user *UserValidator) Validate() map[string]string {
	v := validator.New()
	v.RegisterValidation("password", validPassword)

	errors := make(map[string]string, 0)
	if err := v.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var msg string
			field := string(err.Field())
			switch field {
			case "LastName":
				switch err.ActualTag() {
				case "required":
					msg = fmt.Sprintf("名前を入力してください")
				case "min":
					fallthrough
				case "max":
					msg = fmt.Sprintf("名前は１文字以上25文字以下で入力してください")
				}
			case "FirstName":
				switch err.ActualTag() {
				case "required":
					msg = fmt.Sprintf("名前を入力してください")
				case "min":
					fallthrough
				case "max":
					msg = fmt.Sprintf("名前は１文字以上25文字以下で入力してください")
				}
			case "Email":
				switch err.ActualTag() {
				case "required":
					msg = fmt.Sprintf("メールアドレスを入力してください")
				case "email":
					msg = fmt.Sprintf("メールアドレスの形式で入力してください")
				}
			case "Password":
				switch err.ActualTag() {
				case "required":
					msg = fmt.Sprintf("パスワードを入力してください")
				case "password":
					msg = fmt.Sprintf("パスワードは半角英数字８文字以上で入力してください")
				}
			}
			errors[field] = msg
		}
	}
	return errors
}

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
