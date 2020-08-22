package handler

import (
	"context"
	"grpc-chat/api/application/service"
	"grpc-chat/api/application/validation"
	"grpc-chat/api/gen/pb"
)

func (h *handler) SignUp(
	ctx context.Context,
	req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	lastName := req.GetLastName()
	firstName := req.GetFirstName()
	email := req.GetEmail()
	password := req.GetPassword()

	// 入力バリデーション
	v := validation.NewUserValidator(
		lastName, firstName, email, password)

	errors := v.Validate()
	if len(errors) > 0 {
		return &pb.SignUpResponse{
			Result: false,
			Token:  "",
			Errors: errors,
		}, nil
	}

	s := service.NewService(h.repository)
	res, err := s.SignUp(lastName, firstName, email, password)
	if err != nil {
		return nil, err
	}

	return res, nil
}
