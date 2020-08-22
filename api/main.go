package main

import (
	"grpc-chat/api/application/handler"
	"grpc-chat/api/application/repository"
	"grpc-chat/api/gen/pb"
	"log"
	"net"

	_ "github.com/jinzhu/gorm/dialects/mysql"

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

	r, err := repository.NewRepository()
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.NewHandler(r)
	pb.RegisterAuthServer(s, h)
	reflection.Register(s)

	log.Println("starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
