package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) Say() {
	fmt.Printf("I can say, I am %s, %d years old\n", h.name, h.age)
}

func (h Human) Run() {
	fmt.Printf("%s is running\n", h.name)
}

type Singer struct {
	Human
	collecton string
}

func (s Singer) Sing() {
	fmt.Printf("%s can sing %s\n", s.name, s.collecton)
}

type Student struct {
	Human
	lesson string
}

func (s Student) Learn() {
	fmt.Printf("%s nee learn %s\n", s.name, s.lesson)
}

// interface
type Men interface {
	Say()
	Run()
}

func main() {
	tom := Singer{
		Human:     Human{"Tom", 22},
		collecton: "learn golang",
	}
	peter := Student{
		Human:  Human{"Perter", 18},
		lesson: "learn php",
	}
	jerrk := Human{"jerrk", 26}

	var men Men
	// men 存 tom
	men = tom
	men.Say()
	men.Run()

	men = peter
	men.Say()
	men.Run()

	men = jerrk
	men.Say()
	men.Run()

	x := make([]Men, 3)
	x[0], x[1], x[2] = tom, peter, jerrk

	for _, value := range x {
		value.Say()
		value.Run()
	}
}
