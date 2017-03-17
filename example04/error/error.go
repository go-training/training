package error

import (
	"fmt"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	Title   string
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.Title, e.Message)
}

func IsMyError(err error) bool {
	_, ok := err.(MyError)
	return ok
}
