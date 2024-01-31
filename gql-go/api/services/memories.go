package services

import (
	"context"

	"github.com/4x2Hach1/learning/gql-go/api/db"
	"github.com/4x2Hach1/learning/gql-go/api/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MemoriesService interface {
	// Querys //////////////////////////////
	Memory(ctx context.Context, id int) (*models.Memory, error)
	Memories(ctx context.Context) ([]*models.Memory, error)
	// Mutations //////////////////////////////
	NewMemory(ctx context.Context, input models.NewMemory) (bool, error)
	UpdateMemory(ctx context.Context, input models.UpdateMemory) (bool, error)
	DeleteMemory(ctx context.Context, input models.DeleteMemory) (bool, error)
}

type memoriesService struct {
	*serviceInfra
}

func convertMemory(memory *db.Memory) *models.Memory {
	return &models.Memory{
		ID:       memory.ID,
		UserID:   memory.UsersID,
		Title:    memory.Title,
		Date:     memory.Date,
		Location: memory.Location.Ptr(),
		Detail:   &memory.Detail,
	}
}

func convertMemorySlice(memories []*db.Memory) []*models.Memory {
	results := make([]*models.Memory, 0, len(memories))
	for _, memory := range memories {
		results = append(results, convertMemory(memory))
	}

	return results
}

func (s *memoriesService) Memory(ctx context.Context, id int) (*models.Memory, error) {
	memory, err := db.Memories(qm.Select(), db.MemoryWhere.ID.EQ(id)).One(ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return convertMemory(memory), nil
}

func (s *memoriesService) Memories(ctx context.Context) ([]*models.Memory, error) {
	memories, err := db.Memories(qm.Select()).All(ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return convertMemorySlice(memories), nil
}

func (s *memoriesService) NewMemory(ctx context.Context, input models.NewMemory) (bool, error) {
	memory := db.Memory{
		UsersID:  1,
		Title:    input.Title,
		Date:     input.Date,
		Location: null.StringFromPtr(input.Location),
		Detail:   *input.Detail,
	}

	if err := memory.Insert(ctx, s.exec, boil.Infer()); err != nil {
		return false, err
	}
	return true, nil
}

func (s *memoriesService) UpdateMemory(ctx context.Context, input models.UpdateMemory) (bool, error) {
	memory, err := db.Memories(qm.Select(), db.MemoryWhere.ID.EQ(input.ID), db.MemoryWhere.UsersID.EQ(1)).One(ctx, s.exec)
	if err != nil {
		return false, err
	}

	memory.Title = input.Title
	memory.Date = input.Date
	memory.Location = null.StringFromPtr(input.Location)
	memory.Detail = *input.Detail

	if _, err := memory.Update(ctx, s.exec, boil.Infer()); err != nil {
		return false, err
	}
	return true, nil
}

func (s *memoriesService) DeleteMemory(ctx context.Context, input models.DeleteMemory) (bool, error) {
	if _, err := db.Memories(
		qm.Select(),
		db.MemoryWhere.ID.EQ(input.ID),
		db.MemoryWhere.UsersID.EQ(1),
	).DeleteAll(ctx, s.exec); err != nil {
		return false, err
	}
	return true, nil
}
