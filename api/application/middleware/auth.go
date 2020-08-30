package middleware

import (
	"context"
	"fmt"

	"grpc-chat/api/application/config"
	"grpc-chat/api/pkg/slice"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthFunc func(ctx context.Context, fullMethodName string) (context.Context, error)

// 認証用インターセプタ
func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		newCtx, err := verifyToken(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		res, err := handler(newCtx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

var authorizationFreeMethod = []string{
	"/auth.Auth/SignUp",
	"/auth.Auth/LogIn",
}

// CallされたRPCメソッドについて認証が必要かを判定する
func authorizationRequired(methodName string) bool {
	result, err := slice.Contains(methodName, authorizationFreeMethod)
	if err != nil {
		return false
	}
	return result == false
}

// Authorization Headerから認証用トークンを取得しトークンの検証を行う
// 認証に成功した場合アクセスしたユーザーIDを含んだコンテキストを返す
// トークンの検証が不要な場合は検証をスキップする
func verifyToken(ctx context.Context, methodName string) (context.Context, error) {
	if !authorizationRequired(methodName) {
		return nil, nil
	}

	// Authorization headerから認証トークン取得
	token, err := grpc_auth.AuthFromMD(ctx, "bear")
	if err != nil {
		return nil, status.Error(
			codes.Unauthenticated,
			fmt.Sprintf("cannot get token from header: %v", err),
		)
	}

	// Firebase client取得
	client, err := config.NewFirebaseAuthClient()
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("cannot get firebase client: %v", err),
		)
	}

	// 認証トークン検証
	idToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, status.Error(
			codes.Unauthenticated,
			fmt.Sprintf("token verify failed: %v", err),
		)
	}

	return context.WithValue(ctx, "uid", idToken.UID), nil
}
