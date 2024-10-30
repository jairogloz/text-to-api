package client

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"text-to-api/internal/ports"
)

// repository implements ports.ClientRepository and holds all the required components to
// perform storage operations for users against a postgres database.
// Notice that a client is the same that a Supabase user
type repository struct {
	pool *pgxpool.Pool
}

func NewClientRepository(pool *pgxpool.Pool) (ports.ClientRepository, error) {
	r := &repository{
		pool: pool,
	}

	if r.pool == nil {
		return nil, fmt.Errorf("a connection pool is required")
	}

	return r, nil
}
