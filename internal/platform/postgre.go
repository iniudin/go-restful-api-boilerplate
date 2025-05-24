package platform

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, config *Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.New(ctx, config.DATABASE_URL)
	if err != nil {
		return nil, err
	}

	return poolConfig, nil
}
