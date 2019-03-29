package main

import (
	"fmt"
	"runtime"
	"time"
)

type email struct {
	from    string
	to      string
	current int
}

func (e *email) Current(i int) {
	e.current = i
}

func (e *email) From(s string) {
	e.from = s
}

func (e *email) To(s string) {
	e.to = s
}

func (e *email) Send(i int) {
	if e.current != i {
		fmt.Printf("[Wrong] step %d, from: %s, to: %s\n", i, e.from, e.to)
	} else {
		fmt.Printf("step %d, from: %s, to: %s\n", i, e.from, e.to)
	}
}

func (e email) Current2(i int) email {
	e.current = i

	return e
}

func (e email) From2(s string) email {
	e.from = s

	return e
}

func (e email) To2(s string) email {
	e.to = s

	return e
}

func (e email) Send2(i int) {
	if e.current != i {
		fmt.Printf("[Wrong] step %d, from: %s, to: %s\n", i, e.from, e.to)
	} else {
		fmt.Printf("step %d, from: %s, to: %s\n", i, e.from, e.to)
	}
}

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	e := &email{}
	fmt.Println("==============================")

	for i := 0; i < 10; i++ {
		e.Current(i)
		e.From(fmt.Sprintf("example%02d@gmail.com", i))
		e.To(fmt.Sprintf("example%02d@gmail.com", i+1))
		e.Send(i)
	}

	fmt.Println("==============================")
	fmt.Println("Wrong Result")
	for i := 0; i < 10; i++ {
		go func(i int) {
			e.Current(i)
			e.From(fmt.Sprintf("example%02d@gmail.com", i))
			e.To(fmt.Sprintf("example%02d@gmail.com", i+1))
			e.Send(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("==============================")

	fmt.Println("Fix Result 01")
	for i := 0; i < 10; i++ {
		go func(i int) {
			e := &email{}
			e.Current(i)
			e.From(fmt.Sprintf("example%02d@gmail.com", i))
			e.To(fmt.Sprintf("example%02d@gmail.com", i+1))
			e.Send(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("==============================")

	fmt.Println("Fix Result 02")
	for i := 0; i < 10; i++ {
		go func(i int) {
			e.Current2(i).
				From2(fmt.Sprintf("example%02d@gmail.com", i)).
				To2(fmt.Sprintf("example%02d@gmail.com", i+1)).
				Send2(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("==============================")
}
