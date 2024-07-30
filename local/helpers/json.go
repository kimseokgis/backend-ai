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
