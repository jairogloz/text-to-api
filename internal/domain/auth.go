package domain

// AuthType represents the type of authentication.
type AuthType int

// AuthType constants.
const (
	AuthTypeAPIKey AuthType = iota
	AuthTypeToken
)

// AuthResult represents the result of an authentication.
type AuthResult struct {
	ClientID    string
	Environment RequestEnvironment
}
