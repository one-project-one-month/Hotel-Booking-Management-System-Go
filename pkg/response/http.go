// Package response provides standardized response structures for HTTP and service responses
package response

// HTTPSuccessResponse is a standardized structure for successful HTTP responses.
type HTTPSuccessResponse struct {
	Message string
	Data    interface{}
}

// HTTPErrorResponse is a standardized structure for HTTP error responses.
type HTTPErrorResponse struct {
	Message string
	Error   error
}
