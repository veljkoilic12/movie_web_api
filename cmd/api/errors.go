package main

import (
	"fmt"
	"net/http"
)

// This is a generic helper for logging an error message
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// This method is a generic helper for sending JSON-formatted error messages to the client with a given
// status code.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	// Write the response using the writeJSON() helper. If this happens to return an error then log it
	// and fall back to sending the client an empty response with a 500 internal server error.
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// This method will be used when our app encounters an unexpected problem at runtime. It logs the
// detailed error message, then uses errorResponse() to send a 500 status code and JSON response message
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// This method will be used to send a 404 Not Found status code and JSON response to the client
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// This method will be used to send a 405 Method Not Allowed status code and JSON response to the client
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
