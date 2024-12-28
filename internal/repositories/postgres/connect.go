package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func Connect(ctx context.Context, url string) (pool *pgxpool.Pool, closeConn func(), err error) {

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse database URL: %v", err)
	}
	config.MaxConns = 20               // Maximum connections
	config.MinConns = 5                // Minimum idle connections
	config.MaxConnLifetime = time.Hour // Maximum connection lifetime

	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return pool, pool.Close, nil
}
