package main

import "fmt"

type car struct {
	name  string
	color string
}

func (c *car) SetName01(s string) {
	c.name = s
}

func (c car) SetName02(s string) {
	c.name = s
}

func main() {
	toyota := &car{
		name:  "toyota",
		color: "white",
	}

	fmt.Println(toyota.name)
	toyota.SetName01("foo")
	fmt.Println(toyota.name)
	toyota.SetName02("bar")
	fmt.Println(toyota.name)
	toyota.SetName02("test")
	fmt.Println(toyota.name)
}
