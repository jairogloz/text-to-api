package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"text-to-api/internal/ports"
)

// translator implements ports.Translator and holds all the required components to
// perform translations using the OpenAI API.
type translator struct {
	assistantID string
	client      *openai.Client
	logger      ports.Logger
}

// NewOpenAITranslator creates a new instance of the OpenAI translator.
func NewOpenAITranslator(logger ports.Logger, apiKey string, assistantID string) (ports.Translator, error) {
	client := openai.NewClient(apiKey)

	// Try to list the available models as a simple test
	_, err := client.ListModels(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to validate OpenAI client: %v", err)
	}

	t := &translator{
		assistantID: assistantID,
		client:      client,
		logger:      logger,
	}

	if t.assistantID == "" {
		return nil, fmt.Errorf("assistantID can't be empty")
	}
	if t.logger == nil {
		return nil, fmt.Errorf("logger can't be nil")
	}

	return t, nil
}
