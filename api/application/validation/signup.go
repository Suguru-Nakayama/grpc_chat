package validation

import (
	"github.com/go-playground/validator/v10"
)

type SignUpValidator struct {
	LastName  string `validate:"required,min=1,max=25"`
	FirstName string `validate:"required,min=1,max=25"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,password"`
}

func NewSignUpValidator(lastName, firstName, email, password string) *SignUpValidator {
	return &SignUpValidator{
		LastName:  lastName,
		FirstName: firstName,
		Email:     email,
		Password:  password,
	}
}

func (user *SignUpValidator) Validate() map[string]string {
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
					msg = "名前を入力してください"
				case "min":
					fallthrough
				case "max":
					msg = "名前は１文字以上25文字以下で入力してください"
				}
			case "FirstName":
				switch err.ActualTag() {
				case "required":
					msg = "名前を入力してください"
				case "min":
					fallthrough
				case "max":
					msg = "名前は１文字以上25文字以下で入力してください"
				}
			case "Email":
				switch err.ActualTag() {
				case "required":
					msg = "メールアドレスを入力してください"
				case "email":
					msg = "メールアドレスの形式で入力してください"
				}
			case "Password":
				switch err.ActualTag() {
				case "required":
					msg = "パスワードを入力してください"
				case "password":
					msg = "パスワードは半角英数字８文字以上で入力してください"
				}
			}
			errors[field] = msg
		}
	}
	return errors
}
