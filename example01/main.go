package main

import "fmt"

func main() {
	a := 1
	fmt.Println("一天就學會 Go 語言")
	fmt.Println(HelloWorld())

	if a >= 1 {
		fmt.Println("a >= 1")
	}
}

func HelloWorld() string {
	return fmt.Sprint("一天就學會 Go 語言")
}
