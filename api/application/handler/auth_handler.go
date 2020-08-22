package handler

import (
	"context"
	"grpc-chat/api/application/usecase"
	"grpc-chat/api/gen/pb"
)

type AuthHandler interface {
	SignUp(context.Context, *pb.SignUpRequest) (*pb.SignUpResponse, error)
}
type authHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(au usecase.AuthUseCase) AuthHandler {
	return authHandler{
		authUseCase: au,
	}
}

func (ah authHandler) SignUp(
	ctx context.Context,
	req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	lastName := req.GetLastName()
	firstName := req.GetFirstName()
	email := req.GetEmail()
	password := req.GetPassword()

	res, err := ah.authUseCase.SignUp(lastName, firstName, email, password)
	if err != nil {
		return res, err
	}

	return res, nil
}
