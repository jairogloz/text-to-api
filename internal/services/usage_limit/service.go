package usage_limit

import (
	"fmt"
	"text-to-api/internal/ports"
)

// service provides an implementation of the ports.UsageLimitService interface.
type service struct {
	logger ports.Logger
	repo   ports.UsageLimitRepository
}

// NewUsageLimitService creates and returns a new pointer to service, implementing the ports.UsageLimitService interface.
func NewUsageLimitService(logger ports.Logger, r ports.UsageLimitRepository) (ports.UsageLimitService, error) {
	s := &service{
		logger: logger,
		repo:   r,
	}
	if s.logger == nil {
		return nil, fmt.Errorf("non-nil logger is required to create a new instance of UsageLimitService")
	}
	if s.repo == nil {
		return nil, fmt.Errorf("non-nil usage limit repository is required to create a new instance of UsageLimitService")
	}
	return s, nil
}
