package main

import (
	"grpc-chat/api/application/config"
	"grpc-chat/api/application/handler"
	"grpc-chat/api/application/middleware"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":9090"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			middleware.AuthorizationUnaryServerInterceptor(),
		),
	)

	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatalf("Cannot connect db, %v\n", err)
	}

	h := handler.NewHandler(db)
	h.RegisterPBServer(s)
	reflection.Register(s)

	log.Println("starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
