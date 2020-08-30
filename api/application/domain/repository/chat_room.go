package repository

import "grpc-chat/api/application/domain/model"

type ChatRoomRepository interface {
	Create() (*model.ChatRoom, error)
}
