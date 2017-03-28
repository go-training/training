package main

import (
	"fmt"

	"github.com/go-training/training/example02/hello"
)

func main() {
	fmt.Println("一天就學會 Go 語言")

	hi := hello.HelloWorld("appleboy")
	fmt.Println(hi)
}
