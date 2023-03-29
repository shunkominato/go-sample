package dataloader

// import graph gophers with your other imports
import (
	"context"
	"go-gql-sample/app/ent"
	"log"
	"net/http"

	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)
type Loaders struct {
	UserLoader *dataloader.Loader
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(client *ent.Client) *Loaders {
	// define the data loader
	userReader := &UserReader{Client: client}
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userReader.GetUsers),
	}
	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	log.Print(loadersKey)
	log.Print(ctx.Value(loadersKey))
	return ctx.Value(loadersKey).(*Loaders)
}