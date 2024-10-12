package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"text-to-api/internal/domain"
	"text-to-api/internal/ports"
	"time"
)

// Translate translates the given request to one of the given endpoints.
// Todo: implement final translation
// Todo: try with custom assistant
// Todo: try with fine-tuned model (or fine-tuned model assistant)
func (t translator) Translate(ctx context.Context, prompt string, userID string) (*domain.Translation, error) {

	// Todo: get user threadID if any
	threadID := ""
	//var thread openai.Thread
	//var err error

	if threadID == "" {
		// there's no open thread for the user, so create a new thread and run with the prompt
		run, err := t.client.CreateThreadAndRun(
			ctx,
			openai.CreateThreadAndRunRequest{
				RunRequest: openai.RunRequest{
					AssistantID: t.assistantID,
				},
				Thread: openai.ThreadRequest{
					Messages: []openai.ThreadMessage{
						{
							Role:    openai.ThreadMessageRoleUser,
							Content: prompt,
						},
					},
				},
			},
		)
		if err != nil {
			t.logger.Error(ctx, "Error creating thread and run", "error", err)
			return nil, fmt.Errorf("could not create thread and run: %w", err)
		}
		t.logger.Debug(ctx, "Thread and run created", "threadID", run.ThreadID, "runID", run.ID)

		// Wait for the run to finish
		err = waitForRunCompletion(ctx, t.logger, t.client, run.ThreadID, run.ID, 2*time.Second)
		if err != nil {
			t.logger.Error(ctx, "Error waiting for run completion.", "error", err)
			return nil, fmt.Errorf("could not wait for run completion: %w", err)
		}

		// When run is finished, get the completion
		messageList, err := t.client.ListMessage(ctx, run.ThreadID, domain.Ptr(1),
			domain.Ptr("desc"), nil /*after*/, nil /*before*/, domain.Ptr(run.ID))
		if err != nil {
			t.logger.Error(ctx, "Error listing messages", "error", err)
			return nil, fmt.Errorf("could not list messages: %w", err)
		}

		if len(messageList.Messages) == 0 {
			t.logger.Error(ctx, "No messages found")
			return nil, fmt.Errorf("no messages found")
		}

		message := messageList.Messages[0]
		t.logger.Info(ctx, "Message received", "message", message.Content, "role", message.Role)
	}

	// There seems to be an open thread for the user, so add the prompt to the thread
	// and run
	// Todo: Get thread
	// Todo: Add message to thread
	// Todo: Run thread

	translation := &domain.Translation{
		//Completion: completion,
	}

	return translation, nil

}

func waitForRunCompletion(ctx context.Context, l ports.Logger, client *openai.Client, threadID, runID string, checkInterval time.Duration) error {
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled or timeout reached: %v", ctx.Err())
		case <-ticker.C:
			run, err := client.RetrieveRun(ctx, threadID, runID)
			if err != nil {
				return fmt.Errorf("could not retrieve run: %w", err)
			}
			if run.Status == openai.RunStatusCompleted {
				l.Debug(ctx, "Run completed", "runID", run.ID)
				return nil
			}
			l.Debug(ctx, fmt.Sprintf("Run status: %s", run.Status))
		}
	}
}
