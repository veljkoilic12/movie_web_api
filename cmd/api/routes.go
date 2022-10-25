package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	// initialize a new httprouter router instance from httprouter package
	router := httprouter.New()

	// register the relevant methods, URL patterns and handler functions for our endpoints
	// note that http.MethodGet... are constants which equate to the strings "GET" ...
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	// return the httprouter instance
	return router
}
