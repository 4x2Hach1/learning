package services

import (
	"context"
	"log"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/cache"
	"github.com/4x2Hach1/learning/next-go/api/database/models"
	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type memoryService struct {
	db     *models.Sql
	cache  *cache.Cache
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

func (s *memoryService) MemoryByID(ctx context.Context, p *server.MemoryByIDPayload) (*server.Memory, error) {
	s.logger.Print("server.MemoryByID")
	memory, err := s.db.MemoryByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return memory, nil
}

func (s *memoryService) NewMemory(ctx context.Context, p *server.NewMemoryPayload) (bool, error) {
	s.logger.Print("server.NewMemory")
	token := getUserFromCtx(ctx)
	user, err := s.db.UserById(ctx, token.UserId)
	if err != nil || user == nil {
		return false, err
	}

	date, err := time.Parse("2006-01-02", p.Date)
	if err != nil {
		return false, err
	}

	memory := &models.MemoryModel{
		UsersId:  token.UserId,
		Title:    p.Title,
		Date:     date,
		Location: p.Location,
		Detail:   p.Detail,
	}
	if err := s.db.NewMemory(ctx, memory); err != nil {
		return false, nil
	}
	return true, nil
}

func (s *memoryService) DeleteMemory(ctx context.Context, p *server.DeleteMemoryPayload) (bool, error) {
	s.logger.Print("server.NewMemory")
	token := getUserFromCtx(ctx)
	user, err := s.db.UserById(ctx, token.UserId)
	if err != nil || user == nil {
		return false, err
	}

	if err := s.db.DeleteMemory(ctx, p.ID, *user.ID); err != nil {
		return false, err
	}

	return true, nil
}

func (s *memoryService) UpdateMemory(ctx context.Context, p *server.UpdateMemoryPayload) (bool, error) {
	s.logger.Print("server.UpdateMemory")
	token := getUserFromCtx(ctx)
	user, err := s.db.UserById(ctx, token.UserId)
	if err != nil || user == nil {
		return false, err
	}

	date, err := time.Parse("2006-01-02", p.Date)
	if err != nil {
		return false, err
	}

	memory := &models.MemoryModel{
		ID:       p.ID,
		UsersId:  token.UserId,
		Title:    p.Title,
		Date:     date,
		Location: p.Location,
		Detail:   p.Detail,
	}
	if err := s.db.UpdateMemory(ctx, memory); err != nil {
		return false, err
	}

	return true, nil
}
