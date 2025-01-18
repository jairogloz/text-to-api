package usage_limit

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

// UpsertRequestCountIncrement atomically increments the request count for a given client.
// If no record exists for the client, it inserts a new entry with the initial delta value.
//
// This function ensures that concurrent updates do not cause race conditions by leveraging
// PostgreSQL's `INSERT ... ON CONFLICT ... DO UPDATE` to handle the upsert operation at
// the database level.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellations.
//   - environment: Currently unused, but in the future, it will support environment-specific
//     request tracking (e.g., sandbox vs. production).
//   - clientID: The unique identifier of the client whose request count is being updated.
//   - delta: The number of requests to increment the count by.
//
// Behavior:
// - If the API key exists, the request count is incremented by `delta`.
// - If the API key does not exist, a new record is created with the request count initialized to `delta`.
// - The operation is atomic, preventing race conditions during concurrent updates.
//
// Returns:
// - `nil` on success.
// - An error if the database operation fails.
//
// Future Improvements:
// - Extend the `api_key_usage` table to include an `environment` column for per-environment tracking.
// - Modify the conflict resolution condition to include `environment` (e.g., `ON CONFLICT (api_key, environment)`).
// - Return the updated request count to the caller if needed.
func (r repository) UpsertRequestCountIncrement(ctx context.Context, environment domain.RequestEnvironment, clientID string, delta int) error {
	query := `
		INSERT INTO user_usage (user_id, request_count)
		VALUES ($1, $2)
		ON CONFLICT (user_id) 
		DO UPDATE SET request_count = user_usage.request_count + EXCLUDED.request_count;
	`

	_, err := r.pool.Exec(ctx, query, clientID, delta)
	if err != nil {
		r.logger.Error(ctx, "failed to increment request count", "client_id", clientID, "error", err)
		return fmt.Errorf("failed to increment request count for client %s: %w", clientID, err)
	}

	return nil
}
