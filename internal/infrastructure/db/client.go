package db

import (
	"fmt"
	"go-gql-sample/app/ent"
)

type Database struct {
	client *ent.Client
}

func NewDatabase() (*Database, error) {
	// ここでデータベース接続処理を行います
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"rails_api-db-1", "5432", "postgres", "rails_sample", "postgres"))
	if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &Database{client: client}, nil
}

func (d *Database) Close() error {
	// ここでデータベース接続のクローズ処理を行います
	return d.client.Close()
}