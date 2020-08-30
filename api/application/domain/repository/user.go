package repository

import (
	"grpc-chat/api/application/domain/model"
)

type UserRepository interface {
	Create(lastName, firstName, email, password string) (*model.User, error)
	FindByEmail(email string) *model.User
	FindAllByIds(userIds []uint32) ([]*model.User, error)
}
