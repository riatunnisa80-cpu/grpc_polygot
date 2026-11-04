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
	log.Printf("Menerima request untuk User ID: %d\n", req.GetUserId())
	if req.GetUserId() == 1 {
		return &pb.UserResponse{
			Id:    1,
			Name:  "Budi Santoso",
			Phone: "+6281234567890",
		}, nil
	}
	return &pb.UserResponse{
		Id:    req.GetUserId(),
		Name:  "User Tidak Ditemukan",
		Phone: "-",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("gagal membuka listener: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Server gRPC berjalan di :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gagal menjalankan server: %v", err)
	}
}