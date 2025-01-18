package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// UsageLimitService exposes methods for checking usage limits.
type UsageLimitService interface {
	CheckUsageLimit(ctx context.Context, clientID string, environment domain.RequestEnvironment) error
}

// UsageLimitRepository exposes methods for interacting with the usage limit data in the database.
type UsageLimitRepository interface {
	GetCurrentRequestCount(ctx context.Context, environment domain.RequestEnvironment, clientID string) (int, error)
	ResetRequestCount(environment domain.RequestEnvironment, clientID string) error
	UpsertRequestCountIncrement(ctx context.Context, environment domain.RequestEnvironment, clientId string, delta int) error
}
