package resolver

import (
	"go-gql-sample/app/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Client *ent.Client
	// todos []*model.Todo
	// todoStatus []*model.TodoStatus
}
