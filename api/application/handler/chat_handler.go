package handler

import (
	"context"
	"grpc-chat/api/application/infrastructure/persistence"
	"grpc-chat/api/application/usecase"
	"grpc-chat/api/gen/pb"

	"github.com/jinzhu/gorm"
)

type ChatHandler interface {
	CreateChatRoom(context.Context, *pb.CreateChatRoomRequest) (*pb.CreateChatRoomResponse, error)
}

type chatHandler struct {
	chatUseCase usecase.ChatUseCase
}

func NewChatHandler(db *gorm.DB) ChatHandler {
	up := persistence.NewUserPesistence(db)
	crp := persistence.NewChatRoomPersistence(db)
	cmp := persistence.NewChatMemberPersistence(db)
	return chatHandler{
		chatUseCase: usecase.NewChatUseCase(up, crp, cmp, db),
	}
}

func (ch chatHandler) CreateChatRoom(
	ctx context.Context,
	req *pb.CreateChatRoomRequest,
) (*pb.CreateChatRoomResponse, error) {
	userIds := req.GetUserIds()
	res, err := ch.chatUseCase.CreateChatRoom(ctx, userIds)
	if err != nil {
		return nil, err
	}

	return res, nil
}
