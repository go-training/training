package main

// Comparable constraint

import (
	"constraints"
	"errors"
	"fmt"
)

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

func sum[T constraints.Ordered](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

func main() {
	i, err := indexOfInteger([]int{1, 2, 3}, 3)
	fmt.Println(i, err)

	i, err = indexOfFloat([]float64{1.1, 2.2, 3.3, 4.4}, 4.4)
	fmt.Println(i, err)

	fmt.Println(sum([]int{1, 2, 3, 4}))
	fmt.Println(sum([]float64{1.1, 2.2, 3.3, 4.4}))
	fmt.Println(sum([]float32{1.1, 2.2, 3.3, 4.4, 5.5}))
}
