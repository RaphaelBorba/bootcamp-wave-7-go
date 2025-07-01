package apperrors

import "errors"

var (
	ErrResourceNotExists     = errors.New("resource does not exist")
	ErrInternalError         = errors.New("internal error")
	ErrValidation            = errors.New("validation error")
	ErrResourceAlreadyExists = errors.New("resource already exists")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrForbidden             = errors.New("forbidden")
	ErrNotFound              = errors.New("not found")
	ErrDatabase              = errors.New("database error")
)
