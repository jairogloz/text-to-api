package client

import (
	"fmt"
	"text-to-api/internal/ports"
	postgresPorts "text-to-api/internal/repositories/postgres/ports"
)

// repository implements ports.ClientRepository and holds all the required components to
// perform storage operations for users against a postgres database.
// Notice that a client is the same that a Supabase user
type repository struct {
	logger ports.Logger
	pool   postgresPorts.PgxPool
}

func NewClientRepository(l ports.Logger, pool postgresPorts.PgxPool) (ports.ClientRepository, error) {
	r := &repository{
		logger: l,
		pool:   pool,
	}

	if r.logger == nil {
		return nil, fmt.Errorf("a logger is required")
	}
	if r.pool == nil {
		return nil, fmt.Errorf("a connection pool is required")
	}

	return r, nil
}
