package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestTypeIncorrectFromRequest(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("POST", "/notifications/superType", nil)
	res := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	req = mux.SetURLVars(req, map[string]string{
		"type": "superType",
	})

	PostNotification(res, req)

	assert.Equal(t, 400, res.Code)
	assert.Equal(t, `{"Code":400,"Message":"Couldn't process your request, type not recognized"}`, removeNewLineFromResponse(res.Body.String()))
}

func TestTypeMissingFromRequest(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("POST", "/notifications/superType", nil)
	res := httptest.NewRecorder()

	PostNotification(res, req)

	assert.Equal(t, 400, res.Code)
	assert.Equal(t, `{"Code":400,"Message":"Missing type in the path"}`, removeNewLineFromResponse(res.Body.String()))
}

func removeNewLineFromResponse(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
