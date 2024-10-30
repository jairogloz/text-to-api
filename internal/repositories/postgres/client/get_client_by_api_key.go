package client

import (
	"context"
	"text-to-api/internal/domain"
)

const queryGetByAPIKeyHash = `SELECT 
    ak.created_at AS api_created_at,
    ak.hash,
    ak.environment,
    ak.status,
    ak.user_id,
    u.id,
    u.email,
    u.phone,
    u.created_at AS user_created_at,
    u.last_sign_in_at
FROM 
    api_keys ak
JOIN 
    auth.users u 
ON 
    ak.user_id = u.id
WHERE 
    ak.hash = $1;
    `

func (r repository) GetByAPIKeyHash(ctx context.Context, apiKeyHash string) (client *domain.Client, apiKey *domain.APIKey, err error) {

	client = &domain.Client{}
	apiKey = &domain.APIKey{}

	row := r.pool.QueryRow(ctx, queryGetByAPIKeyHash, apiKeyHash)

	// Scan data into structs
	err = row.Scan(
		&apiKey.CreatedAt,
		&apiKey.Hash,
		&apiKey.Environment,
		&apiKey.Status,
		&apiKey.UserID,
		&client.ID,
		&client.Email,
		&client.Phone,
		&client.CreatedAt,
		&client.LastSignInAt,
	)
	if err != nil {
		// Todo: determine if error is due to no rows found
		return nil, nil, err
	}

	return client, apiKey, nil
}
