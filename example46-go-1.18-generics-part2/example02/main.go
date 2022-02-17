package main

import "fmt"

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

type car struct {
	price int
}

func (c car) String() string {
	return fmt.Sprintf("%d", c.price)
}

func main() {
	val := concat([]stringer{
		car{price: 1},
		car{price: 2},
	})

	fmt.Println(val)
}
