package usage_limit

import (
	"fmt"
	"text-to-api/internal/ports"
	postgresPorts "text-to-api/internal/repositories/postgres/ports"
)

// repository is a PostgresSQL-backed implementation of the ports.UsageLimitRepository interface.
type repository struct {
	logger ports.Logger
	pool   postgresPorts.PgxPool
}

// NewUsageLimitRepository creates a new instance of the repository.
func NewUsageLimitRepository(l ports.Logger, pool postgresPorts.PgxPool) (ports.UsageLimitRepository, error) {
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
