package usecase

import (
	"context"
	"fmt"
	"grpc-chat/api/application/config"
	"grpc-chat/api/application/domain/repository"
	"grpc-chat/api/application/validation"
	"grpc-chat/api/gen/pb"

	fauth "firebase.google.com/go/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthUseCase interface {
	SignUp(lastName, firstName, email, password string) (*pb.SignUpResponse, error)
	LogIn(email, password string) (*pb.LogInResponse, error)
}

type authUseCase struct {
	userRepository repository.UserRepository
}

func NewAuthUseCase(ur repository.UserRepository) AuthUseCase {
	return &authUseCase{
		userRepository: ur,
	}
}

/*
 * サインアップ
 * Firebase AuthenticationとRDBにユーザーを登録する
 */
func (au *authUseCase) SignUp(
	lastName, firstName, email, password string) (*pb.SignUpResponse, error) {
	// 入力バリデーション
	v := validation.NewSignUpValidator(lastName, firstName, email, password)
	errors := v.Validate()
	if len(errors) > 0 {
		return &pb.SignUpResponse{
			Result: false,
			Token:  "",
			Errors: errors,
		}, nil
	}

	// Firebase Client取得
	client, err := config.NewFirebaseAuthClient()
	if err != nil {
		return nil, err
	}

	// Firebase Authenticationにユーザー作成
	params := (&fauth.UserToCreate{}).
		Email(email).
		Password(password)
	u, err := client.CreateUser(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v\n", err)
	}

	// 認証トークン生成
	token, err := client.CustomToken(context.Background(), u.UID)
	if err != nil {
		return nil, fmt.Errorf("error minting custom token: %v\n", err)
	}

	// DBにユーザーを登録
	_, err = au.userRepository.Create(lastName, firstName, email, password)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{
		Result: true,
		Token:  token,
		Errors: nil,
	}, nil
}

/*
 * ログイン
 * メールアドレスとパスワードで認証を行いFirebaseによる認証トークンを生成する
 */
func (au *authUseCase) LogIn(email, password string) (*pb.LogInResponse, error) {
	// 入力バリデーション
	v := validation.NewLogInValidator(email, password)
	errors := v.Validate()
	if len(errors) > 0 {
		return &pb.LogInResponse{
			Result: false,
			Token:  "",
			Errors: errors,
		}, nil
	}

	user := au.userRepository.FindByEmail(email)
	if user == nil {
		st := status.New(codes.NotFound, "No user found for the email address")
		return nil, st.Err()
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// パスワード認証チェック
	if err != nil {
		st := status.New(codes.Unauthenticated, "Authorize failed")
		return nil, st.Err()
	}

	// Firebase Client取得
	client, err := config.NewFirebaseAuthClient()
	if err != nil {
		return nil, err
	}

	// Firebase Authentication上のユーザーを取得
	firebaseUser, err := client.GetUserByEmail(context.Background(), email)
	if err != nil {
		st := status.New(codes.Internal, "Firebase user not found")
		return nil, st.Err()
	}

	// 認証トークン生成
	token, err := client.CustomToken(context.Background(), firebaseUser.UID)
	if err != nil {
		return nil, fmt.Errorf("error minting custom token: %v\n", err)
	}

	return &pb.LogInResponse{
		Result: true,
		Token:  token,
		Errors: nil,
	}, nil
}
