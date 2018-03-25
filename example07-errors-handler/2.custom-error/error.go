package main

import (
	"fmt"
)

// ErrUserNameExist is an error implementation that includes a time and message.
type ErrUserNameExist struct {
	UserName string
}

func (e ErrUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.UserName)
}

// IsErrUserNameExist check error type
func IsErrUserNameExist(err error) bool {
	_, ok := err.(ErrUserNameExist)
	return ok
}
