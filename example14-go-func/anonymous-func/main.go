package main

import "fmt"

func main() {
	foo := func() string {
		return "Hello World1"
	}

	fmt.Println(foo())

	bar := func() {
		fmt.Println("Hello World2")
	}

	bar()

	func() {
		fmt.Println("Hello World3")
	}()

	go func(i, j int) {
		fmt.Println(i + j)
	}(1, 2)
}
