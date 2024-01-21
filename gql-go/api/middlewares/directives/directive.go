package directives

import (
	"context"
	"errors"

	"github.com/4x2Hach1/learning/gql-go/api/execs"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
)

var Directive execs.DirectiveRoot = execs.DirectiveRoot{
	Validation: ValidationInputsValidation,
	IsUserAuth: IsUserAuthenticated,
}

func ValidationInputsValidation(ctx context.Context, obj interface{}, next graphql.Resolver, format *string) (res interface{}, err error) {
	val, err := next(ctx)
	if err != nil {
		return nil, err
	}

	err = validator.New().Var(val, *format)
	if err != nil {
		return nil, err
	} else {
		return val, nil
	}
}

func IsUserAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	token := GetUserFromToken(ctx)
	if token == nil {
		return nil, errors.New("token invalid")
	}

	if token.Token != "token" {
		return nil, errors.New("token invalid")
	}
	return next(ctx)
}
