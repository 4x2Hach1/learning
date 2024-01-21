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

func (s *userService) AuthUser(ctx context.Context, p *server.AuthUserPayload) (*server.User, error) {
	s.logger.Print("server.AuthUser")
	token := getUserFromCtx(ctx)
	return s.db.UserById(ctx, token.UserId)
}

func (s *userService) Users(ctx context.Context, p *server.UsersPayload) ([]*server.User, error) {
	s.logger.Print("server.Users")
	users, err := s.db.AllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) UserByID(ctx context.Context, p *server.UserByIDPayload) (*server.User, error) {
	s.logger.Print("server.UserByID")
	user, err := s.db.UserById(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) NewUser(ctx context.Context, p *server.NewUserPayload) (bool, error) {
	s.logger.Print("server.NewUser")
	user := &models.UserModel{
		Name:     p.Name,
		Email:    p.Email,
		Password: encodePassword(p.Password),
	}
	if err := s.db.NewUser(ctx, user); err != nil {
		return false, err
	}

	return true, nil
}

func (s *userService) UpdateUser(ctx context.Context, p *server.UpdateUserPayload) (bool, error) {
	s.logger.Print("server.UpdateUser")
	token := getUserFromCtx(ctx)

	user := &models.UserModel{ID: token.UserId, Name: p.Name, Email: p.Email}
	if err := s.db.UpdateUser(ctx, user); err != nil {
		return false, err
	}
	return true, nil
}
