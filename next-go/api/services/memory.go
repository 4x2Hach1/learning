package services

import (
	"context"
	"log"

	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type memoryService struct {
	db     *models.Sql
	logger *log.Logger
}

func (s *memoryService) Memories(ctx context.Context, p *server.MemoriesPayload) ([]*server.Memory, error) {
	s.logger.Print("server.Memories")
	memories, err := s.db.AllMemories(ctx)
	if err != nil {
		return nil, err
	}
	return memories, nil
}
