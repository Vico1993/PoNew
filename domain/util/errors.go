package util

type HttpError struct {
	Code    int
	Message string
}

func NewValidationError(message string) *HttpError {
	return &HttpError{
		Code:    400,
		Message: message,
	}
}
