package server

import (
	"context"
	"log"

	pb "user-service/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	log.Printf("Register request for user: %s", req.Username)
	return &pb.UserResponse{Message: "User registered successfully"}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponse, error) {
	log.Printf("Login attempt for user: %s", req.Username)
	return &pb.UserResponse{Message: "User logged in successfully"}, nil
}
