package calc

import (
	"fmt"
)

type Calc struct {
	total int
}

func (c *Calc) Add(i ...int) {
	for _, v := range i {
		c.total += v
	}
}

func (c *Calc) Sub(i ...int) {
	for _, v := range i {
		c.total -= v
	}
}

func (c *Calc) Output() {
	fmt.Println("total:", c.total)
}

func NewCalc(total int) *Calc {
	return &Calc{
		total: total,
	}
}
