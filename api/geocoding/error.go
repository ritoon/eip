package geocoding

import (
	"fmt"
	"net/http"
)

// Error represents a geocoding error.
// It contains an error code, a message, and an origin error wrapped.
type Error struct {
	Code    int
	Message string
	Err     error
}

func newErrorTimeout(message string, err error) *Error {
	return &Error{
		Code:    http.StatusRequestTimeout,
		Message: "geocoding: " + message,
		Err:     err,
	}
}

// ErrIsTimeout returns true if the error is a timeout error.
func ErrIsTimeout(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == http.StatusRequestTimeout
	}
	return false
}

// Error returns the error message.
func (e *Error) Error() string {
	return fmt.Sprintf("db: %v - %s - %v", e.Code, e.Message, e.Err)
}
