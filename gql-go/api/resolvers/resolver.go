package resolvers

import (
	"github.com/4x2Hach1/learning/gql-go/api/middlewares/loaders"
	"github.com/4x2Hach1/learning/gql-go/api/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Srv     services.Services
	Loaders *loaders.Loaders
}
