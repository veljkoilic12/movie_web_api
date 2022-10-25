package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Retrieve the "id" URL parameter from the current request context, then convert it to an integer
// and return it. If the operation is not successful, return 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// when httprouter is parsing a request, any interpolated url parameters will be stored in the
	// request context. We use the ParamsFromContext() function to retrieve a slice containing these
	// parameter names and values
	params := httprouter.ParamsFromContext(r.Context())

	// we can then use the ByName() method to get the value of the "id" parameter from the slice
	// all ID values should be base 10 integers (with a bit size of 64)
	// if the parameter could not be converted, or less than 1, we know the ID is invalid
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Define an envelope type
type envelope map[string]any

// Define a writeJSON() helper for sending responses. This method takes the destination
// http.ResponseWriter, HTTP status code to send, the data to encode to JSON and a header map
// containing any additional HTTP headers we want to include in the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Encode the data to JSON, returning the error if there was one.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')

	// At this point, we know we won't encounter any more errors before writing the response,
	// so it is safe to add any headers that we want to include.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the "Content-Type: application/json" header, then write the status code and JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
