package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/brunoan99/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//	AddUser(client)
	//  AddUserVerbose(client)
	//	AddUsers(client)
	AddUserStreamBoth(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "joazin@example.com",
	}
	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("could not make gRPC request: %v", err)
	}
	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not recive response: %v", err)
		}
		fmt.Println("Status:", stream.Status)
		fmt.Println("User:", stream.User)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "0",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "1",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "2",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "3",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "4",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "5",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	for _, user := range reqs {
		stream.Send(user)
		time.Sleep(time.Second * 3)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}
	fmt.Println("Recived response: ", res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	reqs := []*pb.User{
		{
			Id:    "0",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "1",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "2",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "3",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "4",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
		{
			Id:    "5",
			Name:  "Bruno",
			Email: "brunoan99@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, user := range reqs {
			fmt.Println("Sending User: ", user.Name)
			stream.Send(user)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error receiving response: %v", err)
				break
			}
			fmt.Printf("Received User: %v with status: %v \n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
