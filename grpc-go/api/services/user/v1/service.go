package userv1

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	userv1 "github.com/4x2Hach1/learning/grpc-go/api/gen/proto/user/v1"
	"github.com/4x2Hach1/learning/grpc-go/api/gen/proto/user/v1/userv1connect"
	"github.com/4x2Hach1/learning/grpc-go/api/models"
)

type UserService struct {
	db *models.Sql
}

func NewUserService(db *models.Sql) (string, http.Handler) {
	srv := &UserService{db}
	path, handler := userv1connect.NewUserServiceHandler(srv)

	return path, handler
}

func (s *UserService) User(
	ctx context.Context,
	req *connect.Request[userv1.UserRequest],
) (*connect.Response[userv1.UserResponse], error) {
	log.Println("UserService User headers: ", req.Msg.Id)

	user, err := s.db.UserById(ctx, int(req.Msg.Id))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&userv1.UserResponse{
		User: user,
	})
	res.Header().Set("User-Version", "v1")

	return res, nil
}
