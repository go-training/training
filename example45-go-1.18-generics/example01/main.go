package main

import "fmt"

func show[num int64 | float64](s num) {
	fmt.Println(s)
}

func main() {
	fmt.Println("go 1.18 Generics Example")

	var sum1 int64 = 28
	var sum2 float64 = 29.5

	show(sum1)
	show(sum2)
}
