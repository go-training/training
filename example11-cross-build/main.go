package main

import (
	"fmt"

	"github.com/go-training/training/example11-cross-build/hello"
	"github.com/go-training/training/example11-cross-build/hello2"
	"github.com/go-training/training/example11-cross-build/hello3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello2.Hello())
	fmt.Println(hello3.Hello())
}
