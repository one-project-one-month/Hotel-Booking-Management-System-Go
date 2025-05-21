package response

import "errors"

// ErrNotFound represents an error when a requested entity doesn't exist.
var ErrNotFound = errors.New("entity not found")

// ErrInternalServer represents an internal server error.
var ErrInternalServer = errors.New("internal server error")

// ErrConflict represents an error when an entity operation causes a conflict.
var ErrConflict = errors.New("entity conflict")

// ErrTimeout represents an error when the operation is timeout.
var ErrTimeout = errors.New("operation timeout")

// ServiceResponse is a standardized response structure for internal service operations.
type ServiceResponse struct {
	AppID string      `json:"app_id"`
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}
