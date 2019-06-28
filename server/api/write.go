package api

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteJSON(response http.ResponseWriter, statusCode int, v interface{}) error {
	// set content type to json
	response.Header().Set("Content-Type", "application/json")
	// write the status code
	response.WriteHeader(statusCode)

	// create the json encoder
	encoder := json.NewEncoder(response)
	// display html escape
	encoder.SetEscapeHTML(false)
	// encode the response
	err := encoder.Encode(v)

	return err
}
