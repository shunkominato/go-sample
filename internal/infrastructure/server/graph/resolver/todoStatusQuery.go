package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"go-gql-sample/app/internal/infrastructure/server/graph/model"
	service "go-gql-sample/app/internal/service/todo_status"
)

// TodoStatus is the resolver for the todoStatus field.
func (r *queryResolver) TodoStatus(ctx context.Context, id *int) (*model.TodoStatus, error) {
	res, err := service.GetTodoStatus(ctx, r.client)
	return res, err
}
