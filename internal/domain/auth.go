package domain

// AuthParams represents the parameters required to authenticate a request using an API key.
type AuthParams struct {
	APIKey string
	UserID string
}
