package main

import (
	"fmt"

	"github.com/go-training/training/example13/hello"
	"github.com/go-training/training/example13/hello2"
	"github.com/go-training/training/example13/hello3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello2.Hello())
	fmt.Println(hello3.Hello())
}
