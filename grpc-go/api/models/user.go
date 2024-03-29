package models

import (
	"context"
	"time"

	userv1 "github.com/4x2Hach1/learning/grpc-go/api/gen/proto/user/v1"
)

type UserModel struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func converUser(model *UserModel) *userv1.User {
	return &userv1.User{
		Id:    int32(model.ID),
		Name:  model.Name,
		Email: model.Email,
	}
}

func converUsers(models []*UserModel) []*userv1.User {
	results := make([]*userv1.User, 0, len(models))
	for _, model := range models {
		results = append(results, converUser(model))
	}

	return results
}

func (s *Sql) UserById(ctx context.Context, id int) (*userv1.User, error) {
	user := UserModel{}
	if err := s.db.Get(&user, `SELECT * FROM users WHERE id = ?`, id); err != nil {
		return nil, err
	}

	return converUser(&user), nil
}
