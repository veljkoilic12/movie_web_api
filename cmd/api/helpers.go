package main

import (
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
