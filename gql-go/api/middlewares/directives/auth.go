package directives

import (
	"context"
)

type JwtClaims struct {
	Token string `json:"token"`
}

func DecodeToken(authHeader string) (*JwtClaims, error) {
	return &JwtClaims{Token: authHeader}, nil
}

func GetUserFromToken(ctx context.Context) *JwtClaims {
	switch tokenInfo := ctx.Value(JwtClaims{}).(type) {
	case *JwtClaims:
		return tokenInfo
	default:
		return nil
	}
}
