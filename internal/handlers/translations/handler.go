package translations

import (
	"fmt"
	"text-to-api/internal/ports"
)

// Handler holds the handlers related to translations.
type Handler struct {
	logger  ports.Logger
	service ports.TranslationService
}

// NewTranslationsHandler creates a new Handler for translations requests.
func NewTranslationsHandler(s ports.TranslationService, l ports.Logger) (*Handler, error) {
	h := &Handler{
		logger:  l,
		service: s,
	}

	if h.logger == nil {
		return nil, fmt.Errorf("logger can't be nil")
	}
	if h.service == nil {
		return nil, fmt.Errorf("service can't be nil")
	}

	return h, nil
}
