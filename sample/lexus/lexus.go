package lexus

import ()

// Lexus ....
type Lexus struct {
	Name     string
	Price    float64
	Discount float64
	Total    float64
}

func (t *Lexus) Sum() float64 {
	t.Total = t.Price * t.Discount
	return t.Total
}

func (t *Lexus) SetName(name string) {
	t.Name = name
}

func New(name string) *Lexus {
	car := Lexus{}
	car.Name = name

	return &car
}
