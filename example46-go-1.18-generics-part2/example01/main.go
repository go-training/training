package main

import "fmt"

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func main() {
	fmt.Println(splitAnySlice([]int{1, 2, 3}))
	fmt.Println(splitAnySlice([]int{1, 2, 3, 4}))
	fmt.Println(splitAnySlice([]string{"a", "b", "c", "d"}))
	fmt.Println(splitAnySlice([]float64{1.1, 2.2, 3.3, 4.4}))
}
