package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/4x2Hach1/learning/gql-go/api/models"
)

// NewMemory is the resolver for the newMemory field.
func (r *mutationResolver) NewMemory(ctx context.Context, input models.NewMemory) (bool, error) {
	return r.Srv.NewMemory(ctx, input)
}

// UpdateMemory is the resolver for the updateMemory field.
func (r *mutationResolver) UpdateMemory(ctx context.Context, input models.UpdateMemory) (bool, error) {
	return r.Srv.UpdateMemory(ctx, input)
}

// DeleteMemory is the resolver for the deleteMemory field.
func (r *mutationResolver) DeleteMemory(ctx context.Context, input models.DeleteMemory) (bool, error) {
	return r.Srv.DeleteMemory(ctx, input)
}

// Memory is the resolver for the memory field.
func (r *queryResolver) Memory(ctx context.Context, id int) (*models.Memory, error) {
	return r.Srv.Memory(ctx, id)
}

// Memories is the resolver for the memories field.
func (r *queryResolver) Memories(ctx context.Context) ([]*models.Memory, error) {
	return r.Srv.Memories(ctx)
}
