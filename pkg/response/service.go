package response

import "errors"

var ErrNotFound = errors.New("entity not found")
var ErrInternalServer = errors.New("internal server error")
var ErrConflict = errors.New("entity conflict")

type ServiceResponse struct {
	Data    interface{}
	message string
	Error   error
}
