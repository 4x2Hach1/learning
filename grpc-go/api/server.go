package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/4x2Hach1/learning/grpc-go/api/pb"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedServerServiceServer
}

func (srv *service) Hello(ctx context.Context, p *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}

func auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}

	if token != "auth" {
		return nil, status.Error(codes.Unauthenticated, "token is invalid")
	}

	return ctx, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_auth.UnaryServerInterceptor(auth),
	)))
	pb.RegisterServerServiceServer(s, &service{})
	reflection.Register(s)

	fmt.Println("server is running...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%v", err)
	}
}
