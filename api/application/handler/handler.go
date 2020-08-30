package handler

import (
	"grpc-chat/api/gen/pb"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

type handler struct {
	authHandler AuthHandler
	chatHandler ChatHandler
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{
		authHandler: NewAuthHandler(db),
	}
}

func (h *handler) RegisterPBServer(s *grpc.Server) {
	pb.RegisterAuthServer(s, h.authHandler)
}
