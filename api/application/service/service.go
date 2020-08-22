package service

import "grpc-chat/api/application/repository"

type service struct {
	repository *repository.Repository
}

func NewService(r *repository.Repository) *service {
	return &service{r}
}
