package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

// ToHTTPError returns the appropriate HTTP status code and message for the given error.
func ToHTTPError(err error) (httpStatusCode int, message string) {
	switch {
	case errors.Is(err, domain.ErrorNotFound):
		return fiber.StatusNotFound, fmt.Sprintf("Not found: %s", err.Error())
	case errors.Is(err, domain.ErrorSubscriptionCanceled):
		return fiber.StatusForbidden, "Your subscription is canceled. Access denied."
	case errors.Is(err, domain.ErrorSubscriptionIncomplete):
		return fiber.StatusPaymentRequired, "Subscription is incomplete. Payment required."
	case errors.Is(err, domain.ErrorSubscriptionIncompleteExpired):
		return fiber.StatusForbidden, "Subscription setup expired. Access denied."
	case errors.Is(err, domain.ErrorSubscriptionPastDue):
		return fiber.StatusPaymentRequired, "Subscription is past due. Payment required to continue access."
	case errors.Is(err, domain.ErrorSubscriptionPaused):
		return fiber.StatusForbidden, "Subscription is paused. Access denied."
	case errors.Is(err, domain.ErrorSubscriptionUnpaid):
		return fiber.StatusPaymentRequired, "Subscription is unpaid. Payment required to restore access."
	case errors.Is(err, domain.ErrorValidation):
		return fiber.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error())
	}

	return fiber.StatusInternalServerError, domain.ErrorMessageInternalServerError
}
