package main

import "fmt"

func foo() func() int {
	return func() int {
		return 100
	}
}

func main() {
	bar := foo()
	fmt.Printf("%T\n", bar)
	fmt.Println(bar())

	bar2 := func(i, j float32) float32 {
		return i + j
	}

	fmt.Printf("%T\n", bar2)
	fmt.Println(bar2(1.45, 2.7))
}
