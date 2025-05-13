package response

type HTTPSuccessResponse struct {
	Message string
	Data    interface{}
}

type HTTPErrorResponse struct {
	Message string
	Error   error
}
