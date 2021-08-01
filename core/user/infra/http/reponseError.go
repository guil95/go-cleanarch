package user

type ResponseError struct {
	Message string `json:"message"`
}

func NewResponseError(message string) ResponseError {
	return ResponseError{
		Message: message,
	}
}