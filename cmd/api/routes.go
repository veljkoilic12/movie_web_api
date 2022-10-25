package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	// initialize a new httprouter router instance from httprouter package
	router := httprouter.New()

	// Convert the notFoundResponse() helper to a http.Handler using the http.HandlerFunc() adapter
	// and then set it as the custom error handler for 404 Not Found responses
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// Likewise, convert not allowed response helper to a http.HandlerFunc and set it as the custom
	// error handler for 405 Method Not Allowed responses.
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// register the relevant methods, URL patterns and handler functions for our endpoints
	// note that http.MethodGet... are constants which equate to the strings "GET" ...
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	// return the httprouter instance
	return router
}
