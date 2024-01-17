package services

import (
	"context"
	"log"

	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type userService struct {
	db     *models.Sql
	logger *log.Logger
}

func (s *userService) Users(ctx context.Context, p *server.UsersPayload) ([]*server.User, error) {
	s.logger.Print("server.Users")
	users, err := s.db.AllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
