package main

import "fmt"

type age interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func newGenerics[num age](s1 num) {
	val := float64(s1) + 1
	fmt.Println(val)
}

func total[num age](s1, s2 num) {
	val := float64(s1) + float64(s2)
	fmt.Println(val)
}

func summary[num1, num2 age](s1 num1, s2 num2) {
	val := float64(s1) + float64(s2)
	fmt.Println(val)
}

func main() {
	fmt.Println("go 1.18 Generics Example")

	var sum1 int64 = 28
	var sum2 float64 = 29.5

	newGenerics(sum1)
	newGenerics(sum2)

	var sum3 float64 = 28
	var sum4 float64 = 29.5

	total(sum3, sum4)

	var sum5 int64 = 28
	var sum6 float64 = 29.5

	summary(sum5, sum6)
}
