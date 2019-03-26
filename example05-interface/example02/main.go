package main

import (
	"fmt"
)

type Foo int

func (f Foo) String() string {
	return "foo"
}

func main() {
	var foo Foo = 1000
	fmt.Printf("%d\n", foo)
	fmt.Println(foo)
}
