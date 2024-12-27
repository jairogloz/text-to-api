package api_key

import (
	"fmt"
	"text-to-api/internal/ports"
	postgresPorts "text-to-api/internal/repositories/postgres/ports"
)

// repository implements ports.APIKeyRepository and holds all the required components to
// perform storage operations for API keys against a postgres database.
type repository struct {
	logger ports.Logger
	pool   postgresPorts.PgxPool
}

// NewAPIKeyRepository creates a new repository implementing ports.APIKeyRepository.
func NewAPIKeyRepository(l ports.Logger, pool postgresPorts.PgxPool) (ports.APIKeyRepository, error) {
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
