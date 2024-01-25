package services

import (
	"context"
	"log"

	"github.com/4x2Hach1/learning/next-go/api/cache"
	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type heavyService struct {
	db     *models.Sql
	cache  *cache.Cache
	logger *log.Logger
}

func (s *heavyService) CheckHeavy(ctx context.Context, p *server.CheckHeavyPayload) (int, error) {
	panic("unimplemented")
}

func (s *heavyService) DeleteHeavy(ctx context.Context, p *server.DeleteHeavyPayload) (bool, error) {
	panic("unimplemented")
}

func (s *heavyService) NewHeavy(ctx context.Context, p *server.NewHeavyPayload) (string, error) {
	panic("unimplemented")
}
