package model

type User struct {
	UserId    uint `gorm:"primary_key"`
	LastName  string
	FirstName string
	Email     string
	Password  string
}
