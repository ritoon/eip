package geocoding

import (
	"fmt"
	"net/http"
)

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

func ErrIsTimeout(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == http.StatusRequestTimeout
	}
	return false
}

func (e *Error) Error() string {
	return fmt.Sprintf("db: %v - %s - %v", e.Code, e.Message, e.Err)
}
