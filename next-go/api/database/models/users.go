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
	Password  string    `db:"password" json:"password"`
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

func (s *Sql) UserById(ctx context.Context, id int) (*server.User, error) {
	user := UserModel{}
	if err := s.db.Get(&user, `SELECT * FROM users WHERE id = ?`, id); err != nil {
		return nil, err
	}

	return converUser(&user), nil
}

func (s *Sql) NewUser(ctx context.Context, user *UserModel) error {
	tx := s.db.MustBegin()
	if _, err := tx.NamedExec(
		`INSERT INTO users (name, email, password) VALUES (:name, :email, :password);`,
		user,
	); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *Sql) UpdateUser(ctx context.Context, user *UserModel) error {
	tx := s.db.MustBegin()
	if _, err := tx.NamedExec(
		`UPDATE users SET name = :name, email = :email WHERE id = :id;`,
		user,
	); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *Sql) LoginUser(ctx context.Context, email string, password string) (*server.User, error) {
	user := UserModel{}
	if err := s.db.Get(
		&user, `SELECT * FROM users WHERE email = ? AND password =? Limit 1;`,
		email, password,
	); err != nil {
		return nil, err
	}

	return converUser(&user), nil
}
