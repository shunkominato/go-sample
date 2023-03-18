package repository

import (
	"go-gql-sample/app/ent"
)

type DBClient struct{
	client *ent.Client
}

type Repository interface {
	TodoRepository
}