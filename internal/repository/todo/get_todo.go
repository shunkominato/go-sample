package repository

import (
	"context"
	"errors"
	"go-gql-sample/app/ent"
)

func(r *Repository) GetTodo(ctx context.Context) ([]*ent.Client, error) {
	todo, err := r.client.Todo.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return todo, err
}