package persistence

import (
	"grpc-chat/api/application/domain/model"
	"grpc-chat/api/application/domain/repository"

	"github.com/jinzhu/gorm"
)

type chatRoomPersistence struct {
	db *gorm.DB
}

func NewChatRoomPersistence(db *gorm.DB) repository.ChatRoomRepository {
	return chatRoomPersistence{db}
}

func (crp chatRoomPersistence) Create() (*model.ChatRoom, error) {
	room := new(model.ChatRoom)
	result := crp.db.Table("chat_rooms").Create(room)
	if result.Error != nil {
		return nil, result.Error
	}

	return room, nil
}
