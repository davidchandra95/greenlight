package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

// MarshalJSON returns the json-encoded value for the movie runtime in the format "<runtime> mins"
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// use strconv.Quote function to wrap it in double quotes to  be a valid *JSON string*
	quotedJSONValue := strconv.Quote(jsonValue)

	// convert the quoted string value to a byte slice
	return []byte(quotedJSONValue), nil
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
// Note: because UnmarshalJSON() needs to modify the
// receiver (Runtime type), we must use a pointer receiver for this to work
// correctly. Otherwise, it just only modifying a copy.
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// split the string to isolate the part containing the number.
	parts := strings.Split(unquotedJSONValue, " ")

	// sanity check)
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
