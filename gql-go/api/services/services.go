package services

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type serviceInfra struct {
	exec boil.ContextExecutor
}

type Services interface {
	UsersService
	MemoriesService
}

type services struct {
	*usersService
	*memoriesService
}

func New(exec boil.ContextExecutor) Services {
	infra := &serviceInfra{exec: exec}

	return &services{
		usersService:    &usersService{infra},
		memoriesService: &memoriesService{infra},
	}
}
