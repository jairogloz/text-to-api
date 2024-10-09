package openai

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"text-to-api/internal/ports"
)

// translator implements ports.Translator and holds all the required components to
// perform translations using the OpenAI API.
type translator struct {
	client *openai.Client
}

// NewOpenAITranslator creates a new instance of the OpenAI translator.
func NewOpenAITranslator(apiKey string) (ports.Translator, error) {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	// Try to list the available models as a simple test
	_, err := client.Models.List(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to validate OpenAI client: %v", err)
	}

	t := &translator{
		client: client,
	}

	return t, nil
}
