package translations

import (
	"fmt"
	"text-to-api/internal/ports"
)

// Handler holds the handlers related to translations.
type Handler struct {
	service ports.TranslationService
}

// NewTranslationsHandler creates a new Handler for translations requests.
func NewTranslationsHandler(s ports.TranslationService) (*Handler, error) {
	h := &Handler{
		service: s,
	}

	if h.service == nil {
		return nil, fmt.Errorf("service can't be nil")
	}

	return h, nil
}
