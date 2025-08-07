package main

import (
	"log"
	"net"
	"user-service/internal/server"

	pb "user-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserServer{})

	log.Println("User Service running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
