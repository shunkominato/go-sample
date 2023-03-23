package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	dataloader "go-gql-sample/app/internal/dataloader"
	"go-gql-sample/app/internal/infrastructure/server/graph"
	"go-gql-sample/app/internal/infrastructure/server/graph/model"
)

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj.UserID == nil {
		return nil, nil
	}
	user, err := dataloader.LoadUser(ctx, *obj.UserID)
	if err != nil {
			return nil, err
	}
	return user, nil
}

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
