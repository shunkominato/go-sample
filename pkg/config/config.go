package config

import "os"

func SetConfig() {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		os.Setenv("DB_HOST", "rails_api-db-1")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "postgres")
		os.Setenv("DB_DB_NAME", "rails_sample")
	} else if env == "staging" {
		os.Setenv("DB_HOST", "rails_api-db-1")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "postgres")
		os.Setenv("DB_DB_NAME", "rails_sample")
	} else {
		os.Setenv("DB_HOST", "rails_api-db-1")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "postgres")
		os.Setenv("DB_DB_NAME", "rails_sample")
	}
}