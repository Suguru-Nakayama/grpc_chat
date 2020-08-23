package repository

import (
	"grpc-chat/api/application/domain/model"
)

type UserRepository interface {
	Create(lastName, firstName, uid string) (*model.User, error)
}
