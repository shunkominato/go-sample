package todoService

import (
	"context"
	"errors"
	"go-gql-sample/app/ent"
)

func GetTodo(ctx context.Context, client *ent.Client) ([]*ent.Todo, error) {
	todo, err := client.Todo.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return todo, err
}