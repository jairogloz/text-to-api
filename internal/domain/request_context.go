package domain

import "fmt"

type RequestEnvironment string

const (
	RequestEnvironmentLive    RequestEnvironment = "live"
	RequestEnvironmentSandbox RequestEnvironment = "sandbox"
)

// RequestContext represents the context of a request.
type RequestContext struct {
	ClientID    string
	Environment *RequestEnvironment
	UserID      string
}

// Validate validates the request context.
func (rc *RequestContext) Validate() error {
	if rc.ClientID == "" {
		return fmt.Errorf("clientID is required")
	}
	if rc.Environment == nil {
		return fmt.Errorf("environment is required")
	}
	if rc.UserID == "" {
		return fmt.Errorf("userID is required")
	}
	return nil
}

// NewRequestEnvironment creates a new RequestEnvironment based on the provided environment string.
// It returns an error if the provided environment string is not valid.
//
// Parameters:
//   - env: The environment string, which should be either "live" or "sandbox".
//
// Returns:
//   - The RequestEnvironment if the environment string is valid.
//   - An error if the environment string is invalid.
func NewRequestEnvironment(env string) (RequestEnvironment, error) {
	if env != string(RequestEnvironmentLive) && env != string(RequestEnvironmentSandbox) {
		return "", fmt.Errorf("%w: invalid environment '%s'", ErrorValidation, env)
	}
	e := RequestEnvironment(env)
	return e, nil
}
