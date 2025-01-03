package translations

import (
	"context"
	"errors"
	"fmt"
	"text-to-api/internal/domain"
	"time"
)

// Create creates a new translation based on the user's request.
// It validates the request context and the translation request, retrieves the user from the repository,
// performs the translation, and updates the user metadata if necessary.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - request: The translation request containing the details for the translation.
//   - reqCtx: The request context containing the client ID, user ID, and environment.
//
// Returns:
//   - A pointer to the created domain.Translation.
//   - An error if any issue occurs during the creation process.
func (s service) Create(ctx context.Context, request domain.TranslationRequest, reqCtx domain.RequestContext) (*domain.Translation, error) {

	// Todo: if reqCtx.Environment == "sandbox" use sandbox translation service
	// Todo: handle subscription check

	startValidation := time.Now().UTC()
	if err := reqCtx.Validate(); err != nil {
		s.logger.Error(ctx, "Invalid request context", "error", err)
		return nil, fmt.Errorf("invalid request context: %w", err)
	}
	if err := request.Validate(); err != nil {
		s.logger.Debug(ctx, "Invalid request", "error", err)
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	s.logger.Debug(ctx, "Validation time", "time", time.Since(startValidation).String())

	// Get User from repository
	startGetUser := time.Now().UTC()
	user, err := s.userRepo.GetByID(ctx, *reqCtx.Environment, reqCtx.ClientID, reqCtx.UserID)
	if err != nil && !errors.Is(err, domain.ErrorNotFound) {
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	s.logger.Debug(ctx, "Get user time", "time", time.Since(startGetUser).String())

	var translation domain.Translation
	var newUserMetadata domain.UserMetadata

	switch request.TranslationType {
	case domain.TranslationTypeObject:
		// todo: Update user metadata if required, after translation
		startTranslation := time.Now().UTC()
		var mappedObject interface{}
		mappedObject, newUserMetadata, err = s.translator.TranslateToObject(ctx, request, user)
		if err != nil {
			return nil, fmt.Errorf("could not translate prompt: %w", err)
		}
		s.logger.Debug(ctx, "Translation time", "time", time.Since(startTranslation).String())

		translation = domain.Translation{
			ID:                 "",
			TranslationRequest: request,
			MappedObject:       mappedObject,
		}
	default:
		return nil, fmt.Errorf("%w: translation type '%s' is not supported", domain.ErrorValidation, request.TranslationType)
	}

	// Todo: refactor this function to a separate method
	go func() {
		// If user was nil, we create a new user and save it to the repository
		if user == nil {
			user = &domain.User{
				ClientID: reqCtx.ClientID,
				ID:       reqCtx.UserID,
				Metadata: newUserMetadata,
			}
			err := s.userRepo.Insert(ctx, *reqCtx.Environment, user)
			if err != nil {
				s.logger.Error(ctx, "Error saving user", "error", err)
			}
			return
		}

		// If user was not nil, we update the user metadata only if it was changed
		if !user.Metadata.Equals(newUserMetadata) {
			user.Metadata = newUserMetadata
			err := s.userRepo.Update(ctx, *reqCtx.Environment, user)
			if err != nil {
				s.logger.Error(ctx, "Error updating user", "error", err)
			}
		}
	}()

	return &translation, nil

}
