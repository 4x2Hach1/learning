package loaders

import "github.com/4x2Hach1/learning/gql-go/api/services"

type Loaders struct {
}

func NewLoaders(Srv services.Services) *Loaders {
	return &Loaders{}
}
