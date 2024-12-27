package api_key

import (
	"fmt"
	httpPorts "text-to-api/internal/handlers/ports"
	"text-to-api/internal/ports"
)

// Handler is the receiver for the handler methods related to Customer APIKeys. For
// instance, it is the receiver for the Create APIKey handler.
type Handler struct {
	Logger        ports.Logger
	ReqCtxHandler httpPorts.RequestContextHandler
	Service       ports.APIKeyService
}

// NewAPIKeyHandler creates a new APIKey handler with the provided parameters
func NewAPIKeyHandler(logger ports.Logger, reqCtxHandler httpPorts.RequestContextHandler, apiKeyService ports.APIKeyService) (*Handler, error) {
	h := &Handler{
		Logger:        logger,
		ReqCtxHandler: reqCtxHandler,
		Service:       apiKeyService,
	}
	if h.Logger == nil {
		return nil, fmt.Errorf("nil logger")
	}
	if h.ReqCtxHandler == nil {
		return nil, fmt.Errorf("nil request context handler")
	}
	if h.Service == nil {
		return nil, fmt.Errorf("nil APIKey service")
	}
	return h, nil
}
