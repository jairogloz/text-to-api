package domain

import "errors"

const ErrorMessageInternalServerError = "Internal server error"

var (
	ErrorNotFound   = errors.New("not_found")
	ErrorValidation = errors.New("validation_error")
)
