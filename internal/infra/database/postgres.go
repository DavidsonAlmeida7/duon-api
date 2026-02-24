package database

import (
	"context"
	"duon-api/internal/core/domain/helper"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres() (*pgxpool.Pool, error) {
	//dsn := "postgres://davidson:admin@duon_db:5432/duondb?sslmode=disable"
	//DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	helper.Debug("A", dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return pgxpool.New(ctx, dsn)
}
