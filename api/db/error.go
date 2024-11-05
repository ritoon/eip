package db

import "fmt"

type ErrCode int

const (
	ErrCodeUnknown ErrCode = iota
	ErrCodeNotFound
	ErrCodeInternal
)

func (e ErrCode) String() string {
	switch e {
	case ErrCodeUnknown:
		return "unknown"
	case ErrCodeNotFound:
		return "not found"
	case ErrCodeInternal:
		return "internal"
	}
	return "unknown"
}

type Error struct {
	Code    int `json:"-"`
	Message string
	Err     error
}

func (e *Error) Error() string {
	return fmt.Sprintf("db: %v - %s - %v", e.Code, e.Message, e.Err)
}

func NewError(code ErrCode, message string, err error) *Error {
	return &Error{
		Code:    int(code),
		Message: message,
		Err:     err,
	}
}

func NewErrorNotFound(message string, err error) *Error {
	return NewError(ErrCodeNotFound, message, err)
}

func NewErrorInternal(message string, err error) *Error {
	return NewError(ErrCodeInternal, message, err)
}
