package main

import "fmt"

func main() {
	a := 1
	fmt.Printf(HelloWorld("appleboy"))
	fmt.Println("一天就學會 Go 語言")

	if (a >= 1) { fmt.Println("a >= 1") }
}

func HelloWorld(name string) string {
	return fmt.Sprintf("Hi, %s ", name)
}
