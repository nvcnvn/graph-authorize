package graph

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, idStr string, password string) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, first int, after string) ([]User, error) {
	return []User{
		User{
			ID:    "test_01",
			Roles: []string{"admin"},
		},
		User{
			ID:    "test_01",
			Roles: []string{"user"},
		},
	}, nil
}
func (r *queryResolver) Verify(ctx context.Context, tokenStr string) (bool, error) {
	panic("not implemented")
}
