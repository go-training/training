package main

import (
	"fmt"

	"github.com/go-training/training/sample/toyota"
)

func average(a ...int) (int, int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	average := sum / len(a)

	return sum, average
}

func hello(i int) (a int, b string) {

	if i > 0 {
		a = 1
		b = "one"
		return
	}

	return
}

func main() {
	// var i []int

	// i, s := hello(0)

	// fmt.Println(i, s)

	// _, b := average(1112, 222)
	// c := []int{111, 222, 333, 444}

	// if _, b := average(); b > 0 {
	// 	fmt.Println(b)
	// }

	// fmt.Println(i, a, b)
	// fmt.Println(a)
	// fmt.Println(b)
	// a := 100

	// if ok := hello(0); ok == 1 {
	// 	fmt.Println("ok")
	// }

	// // fmt.Println(i)
	// fmt.Println(a)

	// // ok := hello(0)
	// switch ok := hello(100); ok {
	// case 1:
	// 	fmt.Println("ok")
	// case 0:
	// 	fmt.Println("not ok")
	// }

	t := toyota.Toyota{"car1", 5000, 0.8, 4000}

	fmt.Println(t)
	fmt.Println(t.Sum())
	t.SetName("car2")
	fmt.Println(t.Name)

	d := toyota.New("car03")
	fmt.Println(d)
}
