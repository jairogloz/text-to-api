package domain

import "fmt"

// RequestContext represents the context of a request.
type RequestContext struct {
	ClientID string
	UserID   string
}

// Validate validates the request context.
func (rc *RequestContext) Validate() error {
	if rc.ClientID == "" {
		return fmt.Errorf("clientID is required")
	}
	if rc.UserID == "" {
		return fmt.Errorf("userID is required")
	}
	return nil
}
