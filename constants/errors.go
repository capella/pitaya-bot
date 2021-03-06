package constants

import "errors"

// Errors that can occur during bot execution
var (
	ErrStorageKeyNotFound  = errors.New("storage key not found")
	ErrStorageTypeNotFound = errors.New("storage type not found")
	ErrMalformedObject     = errors.New("malformed object type argument")
)
