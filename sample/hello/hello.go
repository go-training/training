package hello

import (
	"fmt"
)

var total = 100

func init() {
	total += 100
}

func Hello() {
	fmt.Println("Hi")
}
