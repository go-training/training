package main

import (
	"fmt"

	"github.com/go-training/training/example09/lexus"
)

type car interface {
	ShowName() string
	ShowPrice() float64
	ShowDiscount() float64
	ShowColor() string
}

type Toyota struct {
	Name     string
	Price    float64
	Discount float64
	Color    string
}

func (t Toyota) ShowName() string {
	return t.Name
}

func (t Toyota) ShowPrice() float64 {
	return t.Price * t.Discount
}

func (t Toyota) ShowDiscount() float64 {
	return t.Discount
}

func (t Toyota) ShowColor() string {
	return t.Color
}

func detail(c car) {
	fmt.Println("==================")
	fmt.Println("Name:", c.ShowName())
	fmt.Println("Price:", c.ShowPrice())
	fmt.Println("Discount:", c.ShowDiscount())
	fmt.Println("==================")
}

func main() {
	car1 := Toyota{"car1", 3000, 0.8, "white"}
	car2 := Toyota{"car2", 4000, 0.9, "white"}
	car3 := lexus.Lexus{"car3", 5000, 0.7, "blue"}

	detail(car1)
	detail(car2)
	detail(car3)

	car4 := lexus.NewCar("car4", 6000, 0.7, "red")
	detail(car4)
}
