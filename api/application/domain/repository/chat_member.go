package repository

type ChatMemberRepository interface {
	Create(roomId uint32, userIds []uint32) error
}
