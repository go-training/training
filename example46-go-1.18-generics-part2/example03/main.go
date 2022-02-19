package main

// Comparable constraint

import (
	"constraints"
	"errors"
	"fmt"
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

func indexOfInteger[T constraints.Integer](s []T, x T) (int, error) {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func indexOfFloat[T constraints.Float](s []T, x T) (int, error) {
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
	// prints 1 <nil>

	i, err = indexOfInteger([]int{1, 2, 3}, 3)
	fmt.Println(i, err)

	i, err = indexOfFloat([]float64{1.1, 2.2, 3.3, 4.4}, 4.4)
	fmt.Println(i, err)
}
