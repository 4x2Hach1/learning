package services

import (
	"context"
	"log"

	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type userService struct {
	logger *log.Logger
}

func (s *userService) Users(ctx context.Context) (res []*server.User, err error) {
	return nil, nil
}
