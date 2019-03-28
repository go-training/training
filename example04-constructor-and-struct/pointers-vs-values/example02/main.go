package main

import (
	"fmt"
	"time"
)

type email struct {
	from string
	to   string
}

func (e *email) From(s string) {
	e.from = s
}

func (e *email) To(s string) {
	e.to = s
}

func (e *email) send(i int) {
	fmt.Printf("step %d, from: %s, to: %s\n", i, e.from, e.to)
}

func main() {
	e := &email{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			e.From(fmt.Sprintf("example%02d@gmail.com", i))
			e.To(fmt.Sprintf("example%02d@gmail.com", i+1))
			e.send(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
}
