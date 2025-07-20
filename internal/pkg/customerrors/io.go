package customerrors

import "fmt"

// IOError refer to any io related error
type IOError struct {
	Msg     string
	Wrapped error
}

// Error implements the error interface for MyCustomError.
func (e *IOError) Error() string {
	return fmt.Sprintf("[IOError]: %s: %v", e.Msg, e.Wrapped)
}

// Unwrap returns the wrapped error, allowing for error chain inspection.
func (e *IOError) Unwrap() error {
	return e.Wrapped
}

// NewIOError create a new IOError
//
//	msg: word to discribe this error
//	wrapped: original error
func NewIOError(msg string, wrapped error) error {
	return &IOError{
		Msg:     msg,
		Wrapped: wrapped,
	}
}
