package main

import (
	"errors"
	"fmt"

	my "github.com/go-training/training/example04/error"
)

func isEnable(enable bool) (bool, error) {
	if enable {
		return false, errors.New("You can't enable this setting")
	}

	return true, nil
}

func isDisable(disable bool) (bool, error) {
	if disable {
		return false, fmt.Errorf("You can't disable this setting")
	}

	return true, nil
}

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
	_, err := isEnable(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = isDisable(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = isSet(true)
	if err != nil {
		fmt.Println(err.Error())
	}
}
