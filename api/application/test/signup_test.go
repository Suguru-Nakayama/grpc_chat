package main

import (
	"context"
	"grpc-chat/api/gen/pb"
	"grpc-chat/api/model"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	db, err := model.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	app := &Application{db}

	pb.RegisterAuthServer(s, app)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestSignUp(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewAuthClient(conn)

	resp, err := client.SignUp(ctx, &pb.SignUpRequest{
		LastName:  "Tanaka",
		FirstName: "Taro",
		Email:     "test@test.co.jp",
		Password:  "abcd1234",
	})
	if err != nil {
		t.Fatalf("SignUp failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
}
