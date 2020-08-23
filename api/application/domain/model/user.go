package model

type User struct {
	UserId         uint `gorm:"primary_key"`
	LastName       string
	FirstName      string
	FirebaseUserId string
}
