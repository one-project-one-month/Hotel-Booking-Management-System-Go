// Package response provides standardized response structures for HTTP and service responses
package response

// HTTPSuccessResponse is a standardized structure for successful HTTP responses.
type HTTPSuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HTTPErrorResponse is a standardized structure for HTTP error responses.
type HTTPErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
