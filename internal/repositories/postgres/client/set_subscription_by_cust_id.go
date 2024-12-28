package client

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

const setSubscriptionByCustIDQuery = `UPDATE users_data
	SET subscription_id = $2
	WHERE customer_id = $1;`

// SetSubscriptionByCustID updates the subscription ID for a user in the database based on the provided customer ID.
// If the specified customer ID exists in the database, the subscription ID will be updated.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, deadlines, and cancellation signals.
//   - customerID: The unique identifier for the customer in the database.
//   - updateParams: A pointer to a ClientDataUpdate object containing the new subscription ID to be set.
//
// Returns:
//   - error: An error if the operation fails due to database execution issues.
//
// Behavior:
//   - Executes an SQL query to update the subscription ID for the given customer ID.
//   - Logs any errors that occur during the database execution and wraps the error with a descriptive message.
//   - If the operation succeeds, returns nil.
func (r repository) SetSubscriptionByCustID(ctx context.Context, customerID string, updateParams *domain.ClientDataUpdate) error {
	_, err := r.pool.Exec(ctx, setSubscriptionByCustIDQuery, customerID, updateParams.SubscriptionID)
	if err != nil {
		r.logger.Error(ctx, "error upserting user subscription", "error", err)
		return fmt.Errorf("failed to upsert user subscription: %w", err)
	}

	return nil
}
