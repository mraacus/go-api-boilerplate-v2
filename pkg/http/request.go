package httpsuite

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

// ParseRequest parses the incoming HTTP request into a specified struct type,
//
// Validates the parsed request. If the request fails validation or if any error occurs during
// JSON parsing or parameter extraction, it responds with an appropriate HTTP status and error message.
//
// Parameters:
//   - `w`: The `http.ResponseWriter` used to send the response to the client.
//   - `r`: The incoming HTTP request to be parsed.
//
// Returns:
//   - A parsed struct of the specified type `T`, if successful.
//   - An error, if parsing, validation, or parameter extraction fails.
//
// Example usage:
//
//	request, err := ParseRequest[MyRequestType](w, r)
//	if err != nil {
//	    // Handle error
//	}
//
//	// Continue processing the valid request...
func ParseRequest[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var request T
	var empty T
	defer func() { _ = r.Body.Close() }()

	// Decode JSON body if present
	if r.Body != http.NoBody {
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			SendResponse[any](w, http.StatusBadRequest, nil)
			return empty, err
		}
	}

	// Ensure request object is properly initialized
	if isRequestNil(request) {
		request = reflect.New(reflect.TypeOf(request).Elem()).Interface().(T)
	}

	// Validate the request
	if validationErr := IsRequestValid(request); validationErr != nil {
		SendResponse[any](w, http.StatusBadRequest, nil)
		return empty, errors.New("validation error")
	}

	return request, nil
}

// isRequestNil checks if a request object is nil or an uninitialized pointer.
func isRequestNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}