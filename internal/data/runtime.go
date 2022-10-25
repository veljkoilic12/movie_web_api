package data

import (
	"fmt"
	"strconv"
)

// Runtime is a custom type which has the underlying type int32 (the same as our Movie struct field)
type Runtime int32

// Implement a MarshalJSON() method on the Runtime type so that is satisfies the json.Marshaler interface.
// This should return the JSON-encoded value for the movie runtime (string in the format "<runtime> mins")
func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	// Wrap the string in double quotes to. It needs to be surrounded by double quotes in order to be a
	// valid *JSON string*
	quotedJSONValue := strconv.Quote(jsonValue)

	// Convert the quoted string value to a byte slice and return it.
	return []byte(quotedJSONValue), nil
}
