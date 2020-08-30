package validation

import (
	"grpc-chat/api/application/domain/repository"

	"github.com/go-playground/validator/v10"
)

type CreateChatRoomValidator struct {
	UserIds        []uint32 `validate:"dive,required"`
	userRepository repository.UserRepository
}

func NewCreateChatRoomValidator(
	userIds []uint32,
	ur repository.UserRepository,
) *CreateChatRoomValidator {
	return &CreateChatRoomValidator{
		UserIds:        userIds,
		userRepository: ur,
	}
}

func (ccrv *CreateChatRoomValidator) Validate() (map[string]string, error) {
	v := validator.New()
	errors := make(map[string]string, 0)
	if err := v.Struct(ccrv); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var msg string
			field := string(err.Field())
			switch field {
			case "UserIds":
				switch err.ActualTag() {
				case "required":
					msg = "ユーザーIDを入力してください"
				}
			}
			errors[field] = msg
		}
	}
	result, err := isUserRegistered(ccrv.UserIds, ccrv.userRepository)
	if err != nil {
		return nil, err
	}
	if !result {
		errors["UserIds"] = "登録されているユーザーのIDを入力してください"
	}

	return errors, nil
}

// ユーザーIDが既に登録されているユーザーのものであるかを判定する
func isUserRegistered(userIds []uint32, ur repository.UserRepository) (bool, error) {
	users, err := ur.FindAllByIds(userIds)
	if err != nil {
		return false, err
	}
	return len(users) == len(userIds), nil
}
