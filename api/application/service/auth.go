package service

import (
	"fmt"
	"grpc-chat/api/application/repository"
	"grpc-chat/api/auth"
	"grpc-chat/api/gen/pb"
)

func (s *service) SignUp(
	lastName,
	firstName,
	email,
	password string) (*pb.SignUpResponse, error) {
	// ユーザーをDBに登録
	u, err := repository.NewUser(lastName, firstName, email, password)
	if err != nil {
		return nil, err
	}
	u.Create(s.repository.DB)

	// JWTトークン生成
	jwtToken, err := auth.GenerateToken(u.UserId)
	if err != nil {
		return nil, fmt.Errorf("Cannot generate token: %v", err)
	}

	return &pb.SignUpResponse{
		Result: true,
		Token:  jwtToken,
		Errors: nil,
	}, nil
}
