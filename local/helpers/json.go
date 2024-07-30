package helpers

import (
	"encoding/json"
	"net/http"
)

// setJSONHeader sets the Content-Type header to application/json.
// Takes an HTTP response writer.
func setJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// writeStatus sets the HTTP status code for the response.
// Takes an HTTP response writer and the status code.
func writeStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

// encodeJSON encodes the data into JSON format and writes it to the response.
// Takes an HTTP response writer and the data to be encoded.
// Returns an error if encoding fails.
func encodeJSON(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}