package main

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

func indexOf[T comparable](s []T, x T) (int, error) {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func main() {
	i, err := indexOf([]string{"apple", "banana", "pear"}, "banana")
	fmt.Println(i, err)
	i, err = indexOf([]int{1, 2, 3}, 3)
	fmt.Println(i, err)

	fmt.Println(slices.Index([]string{"apple", "banana", "pear"}, "banana"))
}
