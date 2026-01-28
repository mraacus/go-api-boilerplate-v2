package httpsuite

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Response represents the structure of an HTTP response, including a status code, message, and optional body.
// T represents the type of the `Data` field, allowing this structure to be used flexibly across different endpoints.
type Response[T any] struct {
	Data T `json:"data,omitempty"`
}

// SendResponse sends a JSON response to the client, supporting both success and error scenarios.
//
// Parameters:
//   - w: The http.ResponseWriter to send the response.
//   - code: HTTP status code to indicate success or failure.
//   - data: The main payload of the response (only for successful responses).
func SendResponse[T any](w http.ResponseWriter, code int, data T) {

	// Handle error responses
	if code >= 400{
		writeErrorDetail(w, code)
		return
	}

	// Construct and encode the success response
	response := &Response[T]{
		Data: data,
	}

	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(response); err != nil {
		log.Printf("Error writing response: %v", err)
		writeErrorDetail(w, http.StatusInternalServerError)
		return
	}

	// Send the success response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Printf("Failed to write response body (status=%d): %v", code, err)
	}
}

func writeErrorDetail(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.WriteHeader(code)
}