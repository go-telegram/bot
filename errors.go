package bot

import "errors"

var (
	ErrorForbidden       = errors.New("forbidden")
	ErrorBadRequest      = errors.New("bad request")
	ErrorUnauthorized    = errors.New("unauthorized")
	ErrorTooManyRequests = errors.New("too many requests")
	ErrorNotFound        = errors.New("not found")
	ErrorConflict        = errors.New("conflict")
)
