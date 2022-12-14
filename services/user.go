package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/brunoan99/go-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User{
		Id:    "any_id",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "any_id",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "any_id",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("error recieving stream: %s", err)
		}
		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("Adding: ", req.GetName())
	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error recieving stream: %s", err)
		}
		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})
		if err != nil {
			log.Fatalf("error sending stream: %s", err)
		}
	}
}

func NewUserService() *UserService {
	return &UserService{}
}
