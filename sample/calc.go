package main

import (
	"fmt"
)

type calc interface {
	Init()
	Add(...int)
	Sub(...int)
	Output()
}

type calc01 struct {
	total int
}

func (c *calc01) Init() {
	c.total = 100
}

func (c *calc01) Add(input ...int) {
	for _, val := range input {
		c.total += val * 2
	}
}

func (c *calc01) Sub(input ...int) {
	for _, val := range input {
		c.total -= val * 2
	}
}

func (c *calc01) Output() {
	fmt.Println("Output is", c.total)
}

type calc02 struct {
	total int
}

func (c *calc02) Init() {
	c.total = 200
}

func (c *calc02) Add(input ...int) {
	for _, val := range input {
		c.total += val * 3
	}
}

func (c *calc02) Sub(input ...int) {
	for _, val := range input {
		c.total -= val * 4
	}
}

func (c *calc02) Output() {
	fmt.Println("Output is", c.total)
}

func output(c calc) {
	c.Init()
	c.Add(100, 200, 300)
	c.Sub(50, 60)
	c.Output()
}

func main() {
	c1 := new(calc01)
	c2 := new(calc02)
	output(c1)
	output(c2)
}
