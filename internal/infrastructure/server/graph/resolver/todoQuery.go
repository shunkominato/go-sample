package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"
	"go-gql-sample/app/internal/infrastructure/server/graph"
	"go-gql-sample/app/internal/infrastructure/server/graph/model"
	"go-gql-sample/app/internal/repository"
	todoService "go-gql-sample/app/internal/service/todo"
	"strconv"
)

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id *int) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	repo := repository.NewTodoRepository(r.Client)
	s := todoService.NewTodoService(*repo)
	res, err := s.GetTodoList(ctx)
	var todosModel []*model.Todo
	for _, todo := range res {
		todosModel = append(todosModel, &model.Todo{
			ID:           strconv.Itoa(todo.ID),
			Todo:         todo.Todo,
			TodoStatusID: todo.TodoStatusesID,
			User:         todo.UserID,
		})
	}
	return todosModel, err
}

// TodoPagination is the resolver for the todoPagination field.
func (r *queryResolver) TodoPagination(ctx context.Context, limit int, offset *int) (*model.TodoPagination, error) {
	panic(fmt.Errorf("not implemented: TodoPagination - todoPagination"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
