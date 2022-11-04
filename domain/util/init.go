package util

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string
}

// Format and Send JSON error
func SendError(res http.ResponseWriter, err *HttpError) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.WriteHeader(err.Code)

	encodeErr := json.NewEncoder(res).Encode(err)
	if encodeErr != nil {
		panic(encodeErr)
	}
}
