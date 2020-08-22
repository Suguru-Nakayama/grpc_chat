package handler

import "grpc-chat/api/application/repository"

type handler struct {
	repository *repository.Repository
}

func NewHandler(r *repository.Repository) *handler {
	return &handler{r}
}
