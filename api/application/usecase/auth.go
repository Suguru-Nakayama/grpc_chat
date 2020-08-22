package usecase

import (
	"fmt"
	"grpc-chat/api/application/domain/repository"
	"grpc-chat/api/application/validation"
	"grpc-chat/api/auth"
	"grpc-chat/api/gen/pb"
)

type AuthUseCase interface {
	SignUp(lastName, firstName, email, password string) (*pb.SignUpResponse, error)
}

type authUseCase struct {
	userRepository repository.UserRepository
}

func NewAuthUseCase(ur repository.UserRepository) AuthUseCase {
	return &authUseCase{
		userRepository: ur,
	}
}

func (au *authUseCase) SignUp(lastName, firstName, email, password string) (*pb.SignUpResponse, error) {
	v := validation.NewSignUpValidator(lastName, firstName, email, password)
	errors := v.Validate()
	if len(errors) > 0 {
		return &pb.SignUpResponse{
			Result: false,
			Token:  "",
			Errors: errors,
		}, nil
	}

	user, err := au.userRepository.Create(lastName, firstName, email, password)
	if err != nil {
		return nil, err
	}

	// JWTトークン生成
	jwtToken, err := auth.GenerateToken(user.UserId)
	if err != nil {
		return nil, fmt.Errorf("Cannot generate token: %v", err)
	}

	return &pb.SignUpResponse{
		Result: true,
		Token:  jwtToken,
		Errors: nil,
	}, nil
}
