package main

import (
	"fmt"

	"github.com/go-training/training/example02-golang-package/controller"
)

func main() {
	fmt.Println("一天就學會 Go 語言")

	hi := controller.HelloWorld("appleboy")
	fmt.Println(hi)
}
