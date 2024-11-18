package subscription

import (
	"fmt"
	"text-to-api/internal/ports"
	subscriptionPorts "text-to-api/internal/services/subscription/ports"
)

// service implements ports.SubscriptionService and holds the required components to provide
// subscription-related functionality.
type service struct {
	clientRepository ports.ClientRepository
	logger           ports.Logger
	subscriptionHdl  subscriptionPorts.StripeSubscriptionHandler
}

// NewSubscriptionService creates a new SubscriptionService. It returns an error if any of the required
// parameters is nil.
func NewSubscriptionService(cR ports.ClientRepository, l ports.Logger,
	subsHdl subscriptionPorts.StripeSubscriptionHandler) (ports.SubscriptionService, error) {
	s := &service{
		clientRepository: cR,
		logger:           l,
		subscriptionHdl:  subsHdl,
	}
	if s.clientRepository == nil {
		return nil, fmt.Errorf("nil client repository")
	}
	if s.logger == nil {
		return nil, fmt.Errorf("nil logger")
	}
	if s.subscriptionHdl == nil {
		return nil, fmt.Errorf("nil subscription handler")
	}
	return s, nil
}
