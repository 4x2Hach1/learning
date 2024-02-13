package services

import (
	"context"
	"fmt"

	"github.com/4x2Hach1/learning/grpc-go/api/pb"
)

type serviceUser struct {
	*serviceInfr
}

func (srv *serviceUser) User(ctx context.Context, p *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("User")
	user, err := srv.db.UserById(ctx, int(p.Id))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{User: user}, nil
}
