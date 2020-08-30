package model

import "github.com/jinzhu/gorm"

type ChatMember struct {
	ChatMemberId uint64
	ChatRoomId   uint32
	UserId       uint32
	*gorm.Model
}
