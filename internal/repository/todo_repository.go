package repository

import (
	"context"
	"errors"
	"go-gql-sample/app/ent"
)
type TodoRepository struct {
	client *ent.Client
}

func NewTodoRepository(client *ent.Client) *TodoRepository {
	return &TodoRepository{client: client}
}

func (repo *TodoRepository) GetTodoList(ctx context.Context) ([]*ent.Todo, error) {
	todoList, err := repo.client.Todo.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return todoList, err
}