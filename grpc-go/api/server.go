package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/4x2Hach1/learning/grpc-go/api/pb"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedServerServiceServer
}

func (srv *service) Hello(ctx context.Context, p *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServerServiceServer(s, &service{})

	fmt.Println("server is running...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%v", err)
	}
}
