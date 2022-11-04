package util

func NewValidationError(message string) *HttpError {
	return &HttpError{
		Code:    400,
		Message: message,
	}
}
