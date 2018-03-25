package main

import "fmt"

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	a := 1
	b := 2
	a, b = swap(1, 2)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a, b = b, a
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
