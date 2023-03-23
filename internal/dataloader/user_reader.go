package dataloader

import (
	"context"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/internal/infrastructure/server/graph/model"
	"log"

	"github.com/graph-gophers/dataloader"
)

type UserReader struct {
	Client *ent.Client
}

// GetUsers implements a batch function that can retrieve many users by ID,
// for use in a dataloader
func (u *UserReader) GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	log.Print(keys)
	// read all requested users in a single query
	userIDs := make([]string, len(keys))
	for ix, key := range keys {
		userIDs[ix] = key.String()
	}
	userById ,_ := u.Client.User.Query().All(ctx)
	log.Print(userById)

	// return User records into a map by ID
	// userById := map[string]*model.User{}

	// return users in the same order requested
	output := make([]*dataloader.Result, len(keys))
	// for index, userKey := range keys {
	// 	user, ok := userById[userKey.String()]
	// 	if ok {
	// 		output[index] = &dataloader.Result{Data: user, Error: nil}
	// 	} else {
	// 		err := fmt.Errorf("user not found %s", userKey.String())
	// 		output[index] = &dataloader.Result{Data: nil, Error: err}
	// 	}
	// }
	return output
}

func LoadUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.User), nil
}