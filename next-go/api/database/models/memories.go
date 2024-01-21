package models

import (
	"context"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type MemoryModel struct {
	ID        int       `db:"id" json:"id"`
	UsersId   int       `db:"users_id" json:"users_id"`
	Title     string    `db:"title" json:"title"`
	Date      time.Time `db:"date" json:"date"`
	Location  string    `db:"location" json:"location"`
	Detail    string    `db:"detail" json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func converMemory(model *MemoryModel) *server.Memory {
	date := model.Date.Format("2006-01-02")
	return &server.Memory{
		ID:       &model.ID,
		UsersID:  &model.UsersId,
		Title:    &model.Title,
		Date:     &date,
		Location: &model.Location,
		Detail:   &model.Detail,
	}
}

func converMemories(models []*MemoryModel) []*server.Memory {
	results := make([]*server.Memory, 0, len(models))
	for _, model := range models {
		results = append(results, converMemory(model))
	}

	return results
}

func (s *Sql) AllMemories(ctx context.Context) ([]*server.Memory, error) {
	memories := []*MemoryModel{}
	if err := s.db.Select(&memories, `SELECT * FROM memories`); err != nil {
		return nil, err
	}

	return converMemories(memories), nil
}

func (s *Sql) MemoryByID(ctx context.Context, id int) (*server.Memory, error) {
	memory := MemoryModel{}
	if err := s.db.Get(&memory, `SELECT * FROM memories WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return converMemory(&memory), nil
}

func (s *Sql) NewMemory(ctx context.Context, memory *MemoryModel) error {
	tx := s.db.MustBegin()
	if _, err := tx.NamedExec(
		`INSERT INTO memories (users_id, title, date, location, detail)
		VALUES (:users_id, :title, :date, :location, :detail);`,
		memory,
	); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *Sql) DeleteMemory(ctx context.Context, id int, userId int) error {
	tx := s.db.MustBegin()
	if _, err := tx.Exec(
		`DELETE FROM memories WHERE id = ? AND users_id = ?`,
		id, userId,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *Sql) UpdateMemory(ctx context.Context, memory *MemoryModel) error {
	tx := s.db.MustBegin()
	if _, err := tx.NamedExec(
		`UPDATE memories SET title = :title, date = :date, location = :location, detail = :detail 
		WHERE id = :id AND users_id = :users_id;`,
		memory,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
