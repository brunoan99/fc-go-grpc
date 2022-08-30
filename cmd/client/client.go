package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brunoan99/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	start := time.Now()
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	afterConnect := time.Now()
	client := pb.NewUserServiceClient(connection)
	AddUser(client)
	fmt.Println("main, requesition execution time", time.Since(afterConnect))
	fmt.Println("main, total execution time", time.Since(start))
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "joazin@example.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}
