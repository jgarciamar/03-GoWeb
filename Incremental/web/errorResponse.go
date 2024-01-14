package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	ErrInvalidStatusCode = errors.New("the status code provided is invalid ")
)

// The error function takes a writer, a HTTP status code and a message and
// returns a JSON response to the client
func Error(w http.ResponseWriter, code int, message string) {

	// Verify the validity of the status code, how?
	// Maybe checking if the text returned by StatusText is an empty string

	statusText := http.StatusText(code)

	if statusText == "" {
		//The status text is not valid so it shoud return an error maby
		//Ill just set it to 500
		code = http.StatusInternalServerError
	}

	statusCode := strconv.Itoa(code)

	//Create the erroResponse struct

	response := errorResponse{
		Status:  statusCode,
		Message: message,
	}

	//Send the errorResponse JSON to the client

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)

}
