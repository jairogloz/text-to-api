package usage_limit

import (
	"context"
	"errors"
	"fmt"
	"text-to-api/internal/domain"
)

const reqCountLimit = 5

// CheckUsageLimit checks if a client has exceeded their allowed number of requests.
//
// Todo: implement cache for usage count
// Current Limitation (Testing Stage):
//   - Every client is currently limited to 30 requests total across their account.
//   - This restriction is in place because we are in the testing phase and not yet charging users.
//   - In the future, this will change based on the finalized pricing strategy, allowing for flexible limits
//     based on subscription tiers, daily quotas, or pay-per-use models.
//
// Future Improvements:
// - Implement a caching mechanism to reduce database queries for usage checks.
// - Add logic to reset request counts based on time intervals (e.g., daily/monthly limits).
//
// If the request count reaches or exceeds 30, the function returns a 429 Too Many Requests error.
func (s service) CheckUsageLimit(ctx context.Context, clientID string, environment domain.RequestEnvironment) error {
	// Query current usage from the database
	var reqCount int
	var err error
	reqCount, err = s.repo.GetCurrentRequestCount(ctx, environment, clientID)
	if err != nil {
		if errors.Is(err, domain.ErrorNotFound) {
			reqCount = 0
		} else {
			return fmt.Errorf("failed to get current request count: %w", err)
		}
	}
	// Todo: Reset count if interval has passed
	// Check if usage limit has been reached, if so return 429
	if reqCount >= reqCountLimit {
		return fmt.Errorf("%w: This free trial account is limited to %d requests. Contact the developer for additional access", domain.ErrorUsageLimitExceeded, reqCountLimit)
	}

	// Increment usage count
	err = s.repo.UpsertRequestCountIncrement(ctx, environment, clientID, 1)
	if err != nil {
		return fmt.Errorf("failed to increment request count: %w", err)
	}

	return nil
}
