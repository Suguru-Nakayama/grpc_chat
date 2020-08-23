package persistence

import (
	"grpc-chat/api/application/domain/model"
	"grpc-chat/api/application/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPesistence(db *gorm.DB) repository.UserRepository {
	return userPersistence{db}
}

func (up userPersistence) Create(
	lastName, firstName, uid string) (*model.User, error) {

	user := &model.User{
		LastName:       lastName,
		FirstName:      firstName,
		FirebaseUserId: uid,
	}
	up.db.Table("users").Create(user)

	return user, nil
}
