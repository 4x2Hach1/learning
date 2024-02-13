package services

import (
	"context"

	"github.com/4x2Hach1/learning/grpc-go/api/pb"
)

type serviceHello struct {
	*serviceInfr
}

func (srv *serviceHello) Hello(ctx context.Context, p *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}
