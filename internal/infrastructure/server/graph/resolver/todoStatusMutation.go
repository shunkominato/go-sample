package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"
	"go-gql-sample/app/internal/infrastructure/server/graph/model"
)

// CreateTodoStatus is the resolver for the createTodoStatus field.
func (r *mutationResolver) CreateTodoStatus(ctx context.Context, input model.CreateTodoStatusInput) (*model.TodoStatus, error) {
	panic(fmt.Errorf("not implemented: CreateTodoStatus - createTodoStatus"))
}

// UpdateTodoStatus is the resolver for the updateTodoStatus field.
func (r *mutationResolver) UpdateTodoStatus(ctx context.Context, input model.UpdateTodoStatusInput) (*model.TodoStatus, error) {
	panic(fmt.Errorf("not implemented: UpdateTodoStatus - updateTodoStatus"))
}

// DeleteTodoStatus is the resolver for the deleteTodoStatus field.
func (r *mutationResolver) DeleteTodoStatus(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTodoStatus - deleteTodoStatus"))
}
