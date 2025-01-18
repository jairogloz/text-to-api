package usage_limit

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"text-to-api/internal/domain"
)

const QueryGetUserUsageByUserID = `
SELECT u.user_id, u.request_count, u.last_reset
FROM user_usage u
WHERE u.user_id = $1
`

// GetCurrentRequestCount retrieves the number of API requests made by a given client.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellations.
//   - environment: Currently unused, but in the future, it will allow tracking request counts
//     separately for different environments (e.g., production, sandbox).
//   - clientID: The unique identifier of the client whose request count is being retrieved.
//
// Returns:
// - The number of requests the client has made.
// - An error if the query fails or the client does not exist in the usage records.
//
// Notes:
//   - At the moment, all request counts are stored and retrieved without considering
//     different environments. However, this function is structured to support
//     environment-based request tracking in the future.
//   - If no record exists for the given client, it returns a count of 0 and a
//     domain.ErrorNotFound error.
//   - Logs and returns an error if scanning the database row fails.
func (r repository) GetCurrentRequestCount(ctx context.Context, environment domain.RequestEnvironment, clientID string) (int, error) {
	userUsage := &domain.UserUsage{}

	row := r.pool.QueryRow(ctx, QueryGetUserUsageByUserID, clientID)

	err := row.Scan(
		&userUsage.UserID,
		&userUsage.RequestCount,
		&userUsage.LastReset,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("%w: usage for client with id '%s' not found", domain.ErrorNotFound, clientID)
		}
		r.logger.Error(ctx, "error scanning user usage by user id", "error", err)
		return 0, fmt.Errorf("failed to scan user usage by user id: %w", err)
	}

	return userUsage.RequestCount, nil
}
