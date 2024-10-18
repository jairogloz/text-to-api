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
	Environment RequestEnvironment
	UserID      string
}

// Validate validates the request context.
func (rc *RequestContext) Validate() error {
	if rc.ClientID == "" {
		return fmt.Errorf("clientID is required")
	}
	if rc.Environment == "" {
		return fmt.Errorf("environment is required")
	}
	if rc.UserID == "" {
		return fmt.Errorf("userID is required")
	}
	return nil
}
