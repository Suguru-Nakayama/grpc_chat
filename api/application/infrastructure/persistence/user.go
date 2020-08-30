package persistence

import (
	"fmt"
	"grpc-chat/api/application/domain/model"
	"grpc-chat/api/application/domain/repository"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPesistence(db *gorm.DB) repository.UserRepository {
	return userPersistence{db}
}

func (up userPersistence) Create(
	lastName, firstName, email, password string) (*model.User, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Cannot generate token: %v", err)
	}
	user := &model.User{
		LastName:  lastName,
		FirstName: firstName,
		Email:     email,
		Password:  string(hash),
	}
	up.db.Table("users").Create(user)

	return user, nil
}

func (up userPersistence) FindByEmail(email string) *model.User {
	user := &model.User{}
	up.db.Where("email = ?", email).First(user)

	return user
}

func (up userPersistence) FindAllByIds(userIds []uint32) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if result := up.db.Find(users, userIds); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
