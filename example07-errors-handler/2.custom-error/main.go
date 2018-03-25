package main

import (
	"fmt"

	my "github.com/go-training/training/example07-errors-handler/2.custom-error/error"
)

func isSet(disable bool) (bool, error) {
	if disable {
		return false, my.MyError{
			Title:   "Test Title",
			Message: "Test Message",
		}
	}

	return true, nil
}

func main() {
	_, err := isSet(true)
	if err != nil {
		fmt.Println(err.Error())
	}
}
