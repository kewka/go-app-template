package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPool ...
func NewPool(ctx context.Context, cfg *Config) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, cfg.ConnectionURL())
}
