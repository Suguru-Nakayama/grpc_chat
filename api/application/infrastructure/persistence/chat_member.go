package persistence

import (
	"grpc-chat/api/application/domain/model"
	"grpc-chat/api/application/domain/repository"

	"github.com/jinzhu/gorm"
)

type chatMemberPersistence struct {
	db *gorm.DB
}

func NewChatMemberPersistence(db *gorm.DB) repository.ChatMemberRepository {
	return &chatMemberPersistence{db}
}

func (cmp chatMemberPersistence) Create(roomId uint32, userIds []uint32) error {

	members := make([]*model.ChatMember, 0)
	for _, userId := range userIds {
		members = append(members, &model.ChatMember{
			ChatRoomId: roomId,
			UserId:     userId,
		})
	}

	if result := cmp.db.Create(members); result.Error != nil {
		return result.Error
	}
	return nil
}
