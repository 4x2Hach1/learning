package services

import (
	"context"

	"github.com/4x2Hach1/learning/gql-go/api/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UsersService interface {
	// Querys //////////////////////////////
	User(ctx context.Context, id int) (*models.User, error)
	// Mutations //////////////////////////////
	NewUser(ctx context.Context, input models.NewUser) (bool, error)
}

type usersService struct {
	exec boil.ContextExecutor
}

func (services *usersService) User(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}

func (services *usersService) NewUser(ctx context.Context, input models.NewUser) (bool, error) {
	return true, nil
}
