package main

import (
	"fmt"
	// "strconv"

	"github.com/go-training/training/sample/calc"
	_ "github.com/go-training/training/sample/hello"
	"github.com/go-training/training/sample/lexus"
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
	// fmt.Println(hello.Total)
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

	c := calc.NewCalc(100)
	fmt.Println(c)

	c.Add(100, 200)
	c.Sub([]int{100, 1000}...)
	c.Output()

	t1 := toyota.New("toyota01")
	t2 := lexus.New("lexus02")

	Show(t1)
	Show(t2)

	add(100.009)
}

func add(i interface{}) {
	if v, ok := i.(string); ok {
		fmt.Println(v)
	}

	v, ok := i.(string)
	fmt.Println(v, ok)

	// switch i.(type) {
	// case string:
	// 	fmt.Println("it is string")
	// 	_, err := strconv.ParseInt(i 10, 64)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// case int:
	// 	fmt.Println("it is int")
	// 	// integer to string
	// 	a := fmt.Sprintf("%d", i)
	// 	fmt.Println(a)
	// default:
	// 	fmt.Println("not int and string")
	// }
}

func Show(c Car) {
	fmt.Println(c)
}

type Car interface {
	Sum() float64
	SetName(string)
}
