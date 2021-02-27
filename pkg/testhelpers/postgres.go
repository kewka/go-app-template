package testhelpers

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kewka/go-app-template/pkg/postgres"
)

// SetupPostgres ...
func SetupPostgres() *pgxpool.Pool {
	cfg, err := postgres.LoadConfig()
	if err != nil {
		panic(err)
	}
	pool, err := postgres.NewPool(context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	return pool
}
