package main

import "fmt"

type age interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func newGenerics[num age](s num) {
	val := float64(s) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("go 1.18 Generics Example")

	var sum1 int64 = 28
	var sum2 float64 = 29.5

	newGenerics(sum1)
	newGenerics(sum2)
}
