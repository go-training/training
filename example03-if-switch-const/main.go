package main

import "fmt"

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	a = iota
	b
)

func main() {
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
	fmt.Println(sunday)
	fmt.Println(a)
	fmt.Println(b)

	test := true

	if test, test2 := 1, 2; test+test2 < 10 {
		fmt.Println("test + test2 < 10")
	}

	fmt.Println(test)

	switch a := 1; {
	case a >= 0:
		fmt.Println("a is true")
		fallthrough
	case a > 200:
		fmt.Println("a is false")
	}
}
