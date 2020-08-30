package usecase

import (
	"context"
	"fmt"
	"grpc-chat/api/application/domain/repository"
	"grpc-chat/api/application/validation"
	"grpc-chat/api/gen/pb"

	"github.com/jinzhu/gorm"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatUseCase interface {
	CreateChatRoom(ctx context.Context, userIds []uint32) (*pb.CreateChatRoomResponse, error)
}

type chatUseCase struct {
	userRepository       repository.UserRepository
	chatRoomRepository   repository.ChatRoomRepository
	chatMemberRepository repository.ChatMemberRepository
	db                   *gorm.DB
}

func NewChatUseCase(
	ur repository.UserRepository,
	crr repository.ChatRoomRepository,
	cmr repository.ChatMemberRepository,
	db *gorm.DB,
) ChatUseCase {
	return &chatUseCase{
		userRepository:       ur,
		chatRoomRepository:   crr,
		chatMemberRepository: cmr,
		db:                   db,
	}
}

// チャットルームの新規作成とチャットメンバーの登録を行う
func (cuc chatUseCase) CreateChatRoom(ctx context.Context, userIds []uint32) (*pb.CreateChatRoomResponse, error) {
	// 入力バリデーション
	v := validation.NewCreateChatRoomValidator(userIds, cuc.userRepository)
	validationErrors, err := v.Validate()
	if err != nil {
		return nil, err
	}

	if len(validationErrors) > 0 {
		st := status.New(codes.InvalidArgument, "validation error")
		dt, _ := st.WithDetails(&errdetails.BadRequest{
			FieldViolations: validation.ConvertToBadRequestDetails(validationErrors),
		})
		return nil, dt.Err()
	}

	tx := cuc.db.Begin()

	// チャットルーム登録
	room, err := cuc.chatRoomRepository.Create()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("cannot create chat room: %v\n", err)
	}

	// チャットルームメンバー登録
	if err := cuc.chatMemberRepository.Create(room.ChatRoomId, userIds); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("cannot insert chat member: %v\n", err)
	}

	tx.Commit()

	return &pb.CreateChatRoomResponse{
		ChatRoomId:           room.ChatRoomId,
		LastMessageText:      "",
		LastMessageTimestamp: 0,
	}, nil
}
