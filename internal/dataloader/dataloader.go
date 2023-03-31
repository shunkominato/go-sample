package dataloader

// import graph gophers with your other imports
import (
	"context"
	"go-gql-sample/app/ent"
	"log"

	"github.com/gin-gonic/gin"
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

func Middleware(loaders *Loaders) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	ginContext := ctx.Value("GinContextKey")
	log.Print("ginContext**********")
	log.Print(ginContext)
	gc, ok := ginContext.(*gin.Context)

	log.Print("gin**********")
	log.Print(gc)
	log.Print(ok)
	return ctx.Value(loadersKey).(*Loaders)
}