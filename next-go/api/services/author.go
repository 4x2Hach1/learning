package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"github.com/4x2Hach1/learning/next-go/api/gen/server"
	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
)

type authService struct {
	db     *models.Sql
	logger *log.Logger
}

type UserJwtClaims struct {
	UserId int `json:"user_id"`
	Exp    int `json:"exp"`
}

func getUserFromCtx(ctx context.Context) *UserJwtClaims {
	switch tokenInfo := ctx.Value(UserJwtClaims{}).(type) {
	case *UserJwtClaims:
		return tokenInfo
	default:
		return nil
	}
}

func encodePassword(password string) string {
	hashKey := sha256.Sum256(([]byte(password)))
	encodeHashKey := hex.EncodeToString(hashKey[:])
	return encodeHashKey
}

func (s *authService) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	s.logger.Print("server.JWTAuth", token)

	// トークンのパース
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return ctx, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN")), nil
	})
	if err != nil || !parseToken.Valid {
		s.logger.Print("server.JWTAuth", parseToken, err)
		return ctx, errors.New("invalid token")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return ctx, errors.New("invalid token")
	}

	// トークンの有効期限チェック
	exp := int64(claims["exp"].(float64))
	expTime := time.Unix(exp, 0)
	nowTime := time.Now()
	if nowTime.After(expTime) {
		return ctx, errors.New("invalid token")
	}

	// トークンのctxへの埋込み
	userId := int64(claims["user_id"].(float64))
	user := &UserJwtClaims{UserId: int(userId), Exp: int(exp)}
	authCtx := context.WithValue(ctx, UserJwtClaims{}, user)
	return authCtx, nil
}

func (s *authService) Login(ctx context.Context, p *server.LoginPayload) (*server.Loginauth, error) {
	s.logger.Print("server.authService")
	user, err := s.db.LoginUser(ctx, p.Email, encodePassword(p.Password))
	if err != nil {
		return nil, err
	}

	// ユーザーが存在する場合は、トークンの生成
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(os.Getenv("TOKEN")))
	if err != nil {
		return nil, err
	}

	return &server.Loginauth{Token: &accessToken, User: user}, nil
}
