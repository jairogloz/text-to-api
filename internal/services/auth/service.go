package auth

import (
	"fmt"
	"text-to-api/internal/ports"
)

// service implements the ports.AuthService interface
type service struct {
	clientRepo ports.ClientRepository
	jwtSecret  []byte
	logger     ports.Logger
}

// NewAuthService returns a new instance of service implementing the ports.AuthService interface.
func NewAuthService(clientRepo ports.ClientRepository, jwtSecret []byte, l ports.Logger) (ports.AuthService, error) {
	s := &service{
		clientRepo: clientRepo,
		jwtSecret:  jwtSecret,
		logger:     l,
	}

	if s.clientRepo == nil {
		return nil, fmt.Errorf("clientRepo is nil")
	}
	if s.jwtSecret == nil {
		return nil, fmt.Errorf("jwtSecret is nil")
	}
	if s.logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return s, nil
}
