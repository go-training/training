package main

import (
	"fmt"
)

type toyota struct {
	name     string
	price    float64
	discount float64
	color    string
}

func (t *toyota) showName() string {
	return t.name
}

func (t *toyota) getPrice() float64 {
	return t.price * t.discount
}

func (t *toyota) getDiscount() float64 {
	return t.discount
}

func (t *toyota) getColor() string {
	return t.color
}

func (t *toyota) setColor(color string) {
	t.color = color
}

func (t toyota) updateColor(color string) {
	t.color = color
}

func newCar(name string, price float64, discount float64, color string) *toyota {
	return &toyota{
		name:     name,
		price:    price,
		discount: discount,
		color:    color,
	}
}

func main() {
	car := &toyota{
		name:     "car1",
		price:    4000,
		discount: 0.8,
		color:    "white",
	}

	fmt.Println(car.getPrice())
	car.updateColor("blue")
	fmt.Println(car.getColor())
	car.setColor("red")
	fmt.Println(car.getColor())

	car2 := newCar("car2", 6000, 0.7, "orange")
	fmt.Println(car2.getColor())
}
