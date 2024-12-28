package api_key

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

const (
	updatePreviousDataKeysStatuses = `update api_keys set status = $1 where user_id = $2 and environment = $3;`
	insertDataKey                  = `INSERT INTO api_keys (created_at, hash, environment, status, user_id) 
		VALUES ($1,$2, $3, $4, $5);`
)

// SaveAndRevokePrevious performs the following steps atomically:
// 1. Begins a new database transaction.
// 2. Revokes all previous API keys for the user by updating their status.
// 3. Inserts a new API key record.
// 4. Commits the transaction if all operations succeed, otherwise rolls back the transaction.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, deadlines, and cancellation signals.
//   - newAPIKey: The new API key to be saved.
//
// Returns:
//   - error: An error if any of the operations fail.
func (r repository) SaveAndRevokePrevious(ctx context.Context, newAPIKey domain.APIKey) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		r.logger.Error(ctx, "failed to begin transaction", "error", err)
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil {
			if err.Error() != "tx is closed" {
				r.logger.Warn(ctx, "failed to rollback transaction in deferred statement", "error", err)
			}
		}
	}()

	// Revoke all previous api_keys for the user
	_, err = tx.Exec(ctx, updatePreviousDataKeysStatuses, domain.APIKeyStatusRevoked, newAPIKey.UserID, newAPIKey.Environment)
	if err != nil {
		r.logger.Error(ctx, "failed to update data_key status", "error", err)
		return fmt.Errorf("failed to update data_key status: %w", err)
	}

	// Insert the new data_key
	_, err = tx.Exec(ctx, insertDataKey, newAPIKey.CreatedAt, newAPIKey.Hash, newAPIKey.Environment, newAPIKey.Status,
		newAPIKey.UserID)
	if err != nil {
		r.logger.Error(ctx, "failed to insert new data_key", "error", err)
		return fmt.Errorf("failed to insert new data_key: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		r.logger.Error(ctx, "failed to commit transaction", "error", err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
