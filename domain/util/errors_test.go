package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidHttpErrors(t *testing.T) {
	assert.Equal(t, &HttpError{
		Code:    400,
		Message: "Super Error",
	}, NewValidationError("Super Error"))
}
