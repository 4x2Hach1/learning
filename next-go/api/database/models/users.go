package models

import (
	"context"
	"time"

	"github.com/4x2Hach1/learning/next-go/api/gen/server"
)

type UserModel struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func converUser(model *UserModel) *server.User {
	return &server.User{
		ID:    &model.ID,
		Name:  &model.Name,
		Email: &model.Email,
	}
}

func converUsers(models []*UserModel) []*server.User {
	results := make([]*server.User, 0, len(models))
	for _, model := range models {
		results = append(results, converUser(model))
	}

	return results
}

func (s *Sql) AllUsers(ctx context.Context) ([]*server.User, error) {
	users := []*UserModel{}
	if err := s.db.Select(&users, `SELECT * FROM users`); err != nil {
		return nil, err
	}

	return converUsers(users), nil
}
