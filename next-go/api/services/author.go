package services

import (
	"context"
	"errors"
	"fmt"

	"goa.design/goa/v3/security"
)

func (s *serversrvc) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	fmt.Println(token)
	if token != "token" {
		return ctx, errors.New("invalid token")
	}

	return ctx, nil
}
