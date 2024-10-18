package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"text-to-api/internal/domain"
	"text-to-api/internal/ports"
	"time"
)

// TranslateToObject translates the given request to one of the given endpoints.
// It handles the creation of a new thread and run if no existing thread is found for the user,
// or adds a message to an existing thread and runs it if a thread is found.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - translationRequest: The request to be translated.
//   - user: A pointer to the domain.User for whom the translation is being performed.
//
// Returns:
//   - translatedObject: The translated object as an interface{}.
//   - newUserMetadata: The updated user metadata, which includes the thread ID if a new thread was created.
//   - err: An error if any issue occurs during the translation process.
func (t translator) TranslateToObject(ctx context.Context, translationRequest domain.TranslationRequest, user *domain.User) (translatedObject interface{}, newUserMetadata domain.UserMetadata, err error) {

	requestAsJSON, err := json.Marshal(translationRequest)
	if err != nil {
		return nil, nil, fmt.Errorf("could not marshal request: %w", err)
	}

	newUserMetadata = make(domain.UserMetadata)

	// We try to get the user current threadID from the user metadata
	threadID := ""
	if user != nil {
		threadID = user.Metadata.GetString(MetadataFieldThreadID)
		if len(user.Metadata) == 0 {
			user.Metadata = make(domain.UserMetadata)
		}
		newUserMetadata = user.Metadata
	}

	var run openai.Run
	if threadID == "" {
		t.logger.Debug(ctx, "No thread found for user, will create a new one", "userID", user)
		// there's no open thread for the user, so create a new thread and run with the prompt
		run, err = t.client.CreateThreadAndRun(
			ctx,
			openai.CreateThreadAndRunRequest{
				RunRequest: openai.RunRequest{
					AssistantID: t.assistantID,
				},
				Thread: openai.ThreadRequest{
					Messages: []openai.ThreadMessage{
						{
							Role:    openai.ThreadMessageRoleUser,
							Content: string(requestAsJSON),
						},
					},
				},
			},
		)
		if err != nil {
			t.logger.Error(ctx, "Error creating thread and run", "error", err)
			return nil, nil, fmt.Errorf("could not create thread and run: %w", err)
		}
		t.logger.Debug(ctx, "Thread and run created", "threadID", run.ThreadID, "runID", run.ID)

		// As a new thread was created, we append it to the user metadata
		newUserMetadata[MetadataFieldThreadID] = run.ThreadID

	} else {
		t.logger.Debug(ctx, "Thread found for user", "threadID", threadID)

		// There seems to be an open thread for the user, so add the prompt to the thread
		// and run

		// Add message to thread
		// Todo: once CreateRun can receive an openai.RunRequest with additional messages, use that instead
		// and get rid of this CreateMessage part
		_, err = t.client.CreateMessage(ctx, threadID, openai.MessageRequest{
			Role:    string(openai.ThreadMessageRoleUser),
			Content: string(requestAsJSON),
		})
		if err != nil {
			t.logger.Error(ctx, "Error creating message", "error", err)
			return nil, nil, fmt.Errorf("could not create message: %w", err)
		}
		// Run thread
		run, err = t.client.CreateRun(ctx, threadID, openai.RunRequest{
			AssistantID: t.assistantID,
		})
		if err != nil {
			t.logger.Error(ctx, "Error creating run", "error", err)
			return nil, nil, fmt.Errorf("could not create run: %w", err)
		}
	} // end of if/else

	// Todo: register the average completion time to modify the wait time dynamically
	err = waitForRunCompletion(ctx, t.logger, t.client, run.ThreadID, run.ID, 250*time.Millisecond)
	if err != nil {
		t.logger.Error(ctx, "Error waiting for run completion.", "error", err)
		return nil, nil, fmt.Errorf("could not wait for run completion: %w", err)
	}

	// When run is finished, get the completion
	messageList, err := t.client.ListMessage(ctx, run.ThreadID, domain.Ptr(1),
		domain.Ptr("desc"), nil /*after*/, nil /*before*/, domain.Ptr(run.ID))
	if err != nil {
		t.logger.Error(ctx, "Error listing messages", "error", err)
		return nil, nil, fmt.Errorf("could not list messages: %w", err)
	}

	if len(messageList.Messages) == 0 {
		t.logger.Error(ctx, "No messages found")
		return nil, nil, fmt.Errorf("no messages found")
	}

	message := messageList.Messages[0]
	t.logger.Info(ctx, "Message received", "message", message.Content, "role", message.Role)
	resp := ""
	if msgContent := message.Content; len(msgContent) > 0 {
		msg := msgContent[0]
		if msgText := msg.Text; msgText != nil {
			resp = msgText.Value
		}
	}

	var objectMapped map[string]interface{}
	err = json.Unmarshal([]byte(resp), &objectMapped)
	if err != nil {
		t.logger.Error(ctx, "Error unmarshalling response", "error", err)
		return nil, nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	return &objectMapped, newUserMetadata, nil

}

// waitForRunCompletion waits for the completion of a run in the OpenAI client.
// It periodically checks the status of the run and returns when the run is completed
// or if the context is cancelled or times out.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - l: The logger for logging debug information.
//   - client: The OpenAI client used to retrieve the run status.
//   - threadID: The ID of the thread associated with the run.
//   - runID: The ID of the run to wait for completion.
//   - checkInterval: The interval at which to check the run status.
//
// Returns:
//   - An error if the context is cancelled, times out, or if there is an issue retrieving the run status.
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
