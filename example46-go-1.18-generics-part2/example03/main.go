package main

import (
	"errors"
	"fmt"
)

func indexOf[T comparable](s []T, x T) (int, error) {
	for i, v := range s {
		if v == x {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func main() {
	i, err := indexOf([]string{"apple", "banana", "pear"}, "banana")
	fmt.Println(i, err)
	// prints 1 <nil>
}
