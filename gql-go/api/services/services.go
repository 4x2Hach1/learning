package services

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UsersService
}

type services struct {
	*usersService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		usersService: &usersService{exec: exec},
	}
}
