package main

import (
	"fmt"
	"net/http"
)

// declare a handler which writes a json response with information about the app status, env and ver
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed-format JSON response from a string. Notice on using %q to get quoted text.
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Set the "Content-Type: application/json" header on the response. Go defaults to "text/plain"
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON as the HTTP response body.
	w.Write([]byte(js))
}
