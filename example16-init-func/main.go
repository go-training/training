package main

import (
	"fmt"

	_ "github.com/go-training/training/example16-init-func/bar"
	_ "github.com/go-training/training/example16-init-func/foo"
)

var global = convert()

func convert() int {
	return 100
}

func init() {
	global = 0
}

func main() {
	fmt.Println("global is", global)
}
