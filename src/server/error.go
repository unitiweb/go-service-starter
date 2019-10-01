package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// The structure of an error message
type ErrorMessage struct {
	Code string `json:"code"`
	Status int `json:"status"`
	Message string `json:"message"`
	TimeThrown time.Time `json:"timeThrown"`
	Data []string
}

// The outer wrapper structure of an error message
type Error struct {
	Error ErrorMessage `json:"error"`
}

// A slice that holds all the configured error messages
var apiErrors []ErrorMessage

// Handles the error, formats, and outputs to the response
func HandleError(w http.ResponseWriter, e ErrorMessage) {
	// Set the current timestamp
	e.TimeThrown = time.Now()

	// Convert into Json
	b, err := json.Marshal(Error{Error: e})
	if err != nil {
		io.WriteString(w, "Error could not be converted to json")
	}

	// Write out the error
	w.WriteHeader(e.Status)
	io.WriteString(w, string(b))
}

// Add an error message type to the server
func AddError(c string, s int, m string) {
	e := ErrorMessage{Code: c, Status: s, Message: m}
	apiErrors = append(apiErrors, e)
}

// Find a configured error message
// If none exists then return an unknown error message
func FindError(c string) ErrorMessage {
	// Attempt to find the error by error code
	for _, err := range apiErrors {
		if err.Code == c {
			err.TimeThrown = time.Now()
			return err
		}
	}

	// Since error could not be found thrown an unknown error
	return ErrorMessage{
		Code:       "UnknownError",
		Status:     500,
		Message:    c,
		TimeThrown: time.Time{},
	}
}

// Used to throw and error message by error message code
// It will return an error with the code as the message
func ThrowError(code string) error {
	return errors.New(code)
}
