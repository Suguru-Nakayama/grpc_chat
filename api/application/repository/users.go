package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    uint `gorm:"primary_key"`
	LastName  string
	FirstName string
	Email     string
	Password  string
}

func NewUser(
	lastName string,
	firstName string,
	email string,
	password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Cannot generate token: %v", err)
	}

	user := &User{
		LastName:  lastName,
		FirstName: firstName,
		Email:     email,
		Password:  string(hash),
	}

	return user, nil
}

func (user *User) Create(db *gorm.DB) *User {
	db.Table("users").Create(user)
	return user
}
