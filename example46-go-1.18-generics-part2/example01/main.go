package main

import (
	"errors"
	"fmt"
)

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

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
	fmt.Println(splitAnySlice([]int{1, 2, 3}))
	fmt.Println(splitAnySlice([]int{1, 2, 3, 4}))
	fmt.Println(splitAnySlice([]string{"a", "b", "c", "d"}))
	fmt.Println(splitAnySlice([]float64{1.1, 2.2, 3.3, 4.4}))

	i, err := indexOf([]string{"apple", "banana", "pear"}, "banana")
	fmt.Println(i, err)
	i, err = indexOf([]int{1, 2, 3}, 3)
	fmt.Println(i, err)
	// prints 1 <nil>
}
