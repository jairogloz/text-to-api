package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

// ToHTTPError returns the appropriate HTTP status code and message for the given error.
func ToHTTPError(err error) (httpStatusCode int, message string) {
	if errors.Is(err, domain.ErrorValidation) {
		return fiber.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error())
	}
	return fiber.StatusInternalServerError, "Internal server error"
}
