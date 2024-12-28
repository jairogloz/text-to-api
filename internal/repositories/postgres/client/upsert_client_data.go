package client

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

const upsertClientDataQuery = `INSERT INTO users_data (user_id, subscription_id, customer_id)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id)
	DO UPDATE SET 
	    subscription_id = EXCLUDED.subscription_id,
	    customer_id = EXCLUDED.customer_id;`

// UpsertClientData inserts or updates client data in the database for the given user ID.
// If a row with the specified user ID already exists, the function updates the relevant fields
// (such as subscription ID and customer ID). If no row exists, a new row is inserted.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, deadlines, and cancellation signals.
//   - id: The unique identifier for the user whose data is being upserted.
//   - updateParams: A pointer to a ClientDataUpdate struct containing fields to be updated or inserted.
//
// Returns:
//   - error: An error if the operation fails due to database execution issues.
func (r repository) UpsertClientData(ctx context.Context, id string, updateParams *domain.ClientDataUpdate) error {
	_, err := r.pool.Exec(context.Background(), upsertClientDataQuery, id, updateParams.SubscriptionID,
		updateParams.CustomerID)
	if err != nil {
		r.logger.Error(ctx, "error upserting user data", "error", err)
		return fmt.Errorf("failed to upsert user data: %w", err)
	}

	return nil
}
