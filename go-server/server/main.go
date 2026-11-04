
package main

import (
	"context"
	"log"
	"net"

	pb "go-server/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserProfile(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Name:  "Ria",
		Email: "ria@example.com",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Server gRPC berjalan di port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Server gagal: %v", err)
	}
}
