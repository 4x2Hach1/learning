package services

import (
	"context"

	"github.com/4x2Hach1/learning/gql-go/api/db"
	"github.com/4x2Hach1/learning/gql-go/api/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UsersService interface {
	// Querys //////////////////////////////
	User(ctx context.Context, id int) (*models.User, error)
	// Mutations //////////////////////////////
	NewUser(ctx context.Context, input models.NewUser) (bool, error)
	UpdateUser(ctx context.Context, input models.UpdateUser) (bool, error)
}

type usersService struct {
	*serviceInfra
}

func convertUser(user *db.User) *models.User {
	return &models.User{
		ID:    user.ID,
		Name:  user.Name.String,
		Email: user.Email.String,
	}
}

func (s *usersService) User(ctx context.Context, id int) (*models.User, error) {
	user, err := db.Users(
		qm.Select(),
		db.UserWhere.ID.EQ(id),
	).One(ctx, s.exec)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func (s *usersService) NewUser(ctx context.Context, input models.NewUser) (bool, error) {
	count, err := db.Users(
		qm.Select(),
		db.UserWhere.Email.EQ(null.StringFrom(input.Email)),
	).Count(ctx, s.exec)
	if err != nil || 0 < count {
		return false, err
	}

	user := db.User{
		Name:     null.StringFrom(input.Name),
		Email:    null.StringFrom(input.Email),
		Password: input.Password,
	}

	if err := user.Insert(ctx, s.exec, boil.Infer()); err != nil {
		return false, err
	}

	return true, nil
}

func (s *usersService) UpdateUser(ctx context.Context, input models.UpdateUser) (bool, error) {
	user, err := db.Users(
		qm.Select(),
		db.UserWhere.ID.EQ(input.ID),
	).One(ctx, s.exec)
	if err != nil {
		return false, err
	}

	user.Name = null.StringFrom(input.Name)
	user.Email = null.StringFrom(input.Email)
	user.Password = input.Password
	if _, err := user.Update(ctx, s.exec, boil.Infer()); err != nil {
		return false, err
	}

	return true, nil
}
