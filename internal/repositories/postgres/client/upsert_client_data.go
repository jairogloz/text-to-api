package client

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

const upsertClientDataQuery = `INSERT INTO users_data (user_id, customer_id)
	VALUES ($1, $2)
	ON CONFLICT (user_id)
	DO UPDATE SET customer_id = EXCLUDED.customer_id;`

// UpsertClientData inserts or updates a clients (user in supabase) data in the users_data table.
// If a row with the specified user ID already exists, the function updates the customer_id field;
// otherwise, it inserts a new row. This function is commonly referred to as an "upsert" operation.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, deadlines, and cancellation signals.
//   - id: The unique identifier for the user whose data is being upserted.
//   - updateParams: A pointer to ClientDataUpdate containing fields to be updated or inserted,
//     specifically the CustomerID field in this case.
//
// Returns:
// - error: An error if the upsert operation fails due to a database or execution issue.
//
// Behavior:
//   - If an error occurs during the execution of the upsert, it logs the error with contextual details
//     and returns a wrapped error indicating failure. If successful, it returns nil.
func (r repository) UpsertClientData(ctx context.Context, id string, updateParams *domain.ClientDataUpdate) error {
	_, err := r.pool.Exec(context.Background(), upsertClientDataQuery, id, updateParams.CustomerID)
	if err != nil {
		r.logger.Error(ctx, "error upserting user data", "error", err)
		return fmt.Errorf("failed to upsert user data: %w", err)
	}

	return nil
}
