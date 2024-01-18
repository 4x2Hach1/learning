package services

import (
	"context"
	"errors"
	"log"

	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"goa.design/goa/v3/security"
)

type authService struct {
	db     *models.Sql
	logger *log.Logger
}

func (s *authService) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	s.logger.Print("server.JWTAuth", token)
	if token != "token" {
		return ctx, errors.New("invalid token")
	}

	return ctx, nil
}
