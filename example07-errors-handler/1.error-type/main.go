package main

import (
	"errors"
	"fmt"
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

func main() {
	_, err := isEnable(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = isDisable(true)
	if err != nil {
		fmt.Println(err.Error())
	}
}
