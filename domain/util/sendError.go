package util

import (
	"encoding/json"
	"net/http"
)

// Format and Send JSON error
// TODO: Find a way to properly test this method. With http package be able to assert the http call?
func SendError(res http.ResponseWriter, err *HttpError) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.WriteHeader(err.Code)

	encodeErr := json.NewEncoder(res).Encode(err)
	if encodeErr != nil {
		panic(encodeErr)
	}
}
