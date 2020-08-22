package main

import (
	"grpc-chat/api/application/config"
	"grpc-chat/api/application/handler"
	"grpc-chat/api/application/infrastructure/persistence"
	"grpc-chat/api/application/usecase"
	"grpc-chat/api/gen/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":9090"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatalf("Cannot connect db, %v\n", err)
	}
	up := persistence.NewUserPesistence(db)
	au := usecase.NewAuthUseCase(up)
	ah := handler.NewAuthHandler(au)

	pb.RegisterAuthServer(s, ah)
	reflection.Register(s)

	log.Println("starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
