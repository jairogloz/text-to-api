package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"text-to-api/internal/domain"
)

const queryGetClientByID = `SELECT 
    u.id,
    u.email,
    u.phone,
    u.created_at AS user_created_at,
    u.last_sign_in_at,
    ud.subscription_id,
    ud.customer_id
FROM 
    auth.users u
LEFT OUTER JOIN 
    users_data ud
ON 
    ud.user_id = u.id
WHERE 
    u.id = $1;`

// GetByID retrieves a client record from the database using the provided client ID.
// It queries the database with the client ID and scans the result into a Client struct.
// If no matching record is found, it returns a not found error. If a different error occurs during the
// query, it logs the error and returns it.
//
// Parameters:
//   - ctx: The context for request-scoped values, cancellations, and deadlines.
//   - clientID: The unique identifier for the client in the database.
//
// Returns:
//   - *domain.Client: A pointer to the Client struct populated with the client’s data if found.
//   - error: An error if the client is not found (ErrNotFound) or if any other database error occurs.
func (r repository) GetByID(ctx context.Context, clientID string) (*domain.Client, error) {
	client := &domain.Client{}

	row := r.pool.QueryRow(ctx, queryGetClientByID, clientID)

	err := row.Scan(
		&client.ID,
		&client.Email,
		&client.Phone,
		&client.CreatedAt,
		&client.LastSignInAt,
		&client.SubscriptionID,
		&client.CustomerID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: client with id '%s' not found", domain.ErrorNotFound, clientID)
		}
		r.logger.Error(ctx, "error scanning client by id", "error", err)
	}

	return client, nil
}
