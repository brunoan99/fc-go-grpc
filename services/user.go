package services

import (
	"context"
	"fmt"

	"github.com/brunoan99/go-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "any_id",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
