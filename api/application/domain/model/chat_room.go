package model

import "github.com/jinzhu/gorm"

type ChatRoom struct {
	ChatRoomId uint32 `gorm:"primary_key"`
	*gorm.Model
}
