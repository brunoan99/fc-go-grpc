package main

import (
	"fmt"
	"log"
	"net"

	"github.com/brunoan99/go-grpc/pb"
	"github.com/brunoan99/go-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: $%v", err)
	}

	fmt.Println("Listen tcp on localhost:50051")

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: $%v", err)
	}
}
