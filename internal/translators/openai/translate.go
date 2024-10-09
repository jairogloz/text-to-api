package openai

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"text-to-api/internal/domain"
)

// Translate translates the given request to one of the given endpoints.
// Todo: implement actual functionality
func (t translator) Translate(ctx context.Context, request domain.TranslationRequest, availableEndpoints []domain.Endpoint) (*domain.Translation, error) {

	completion, err := t.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("what is the purpose of life"),
			}),
			Model: openai.F(openai.ChatModelGPT3_5Turbo),
		})
	if err != nil {
		return nil, fmt.Errorf("could not get completion: %w", err)
	}

	translation := &domain.Translation{
		Completion: completion,
	}

	return translation, nil

}
