package domain

import "errors"

const ErrorMessageInternalServerError = "Internal server error"

var (
	ErrorNotFound                      = errors.New("not_found")
	ErrorSubscriptionCanceled          = errors.New("subscription_canceled")
	ErrorSubscriptionIncomplete        = errors.New("subscription_status_incomplete")
	ErrorSubscriptionIncompleteExpired = errors.New("subscription_incomplete_expired")
	ErrorSubscriptionPastDue           = errors.New("subscription_past_due")
	ErrorSubscriptionPaused            = errors.New("subscription_paused")
	ErrorSubscriptionUnpaid            = errors.New("subscription_unpaid")
	ErrorValidation                    = errors.New("validation_error")
)
