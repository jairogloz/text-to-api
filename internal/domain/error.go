package domain

import "errors"

var (
	ErrorNotFound   = errors.New("not_found")
	ErrorValidation = errors.New("validation_error")
)
