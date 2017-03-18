package toyota

import ()

// Toyota ....
type Toyota struct {
	Name     string
	Price    float64
	Discount float64
	Total    float64
}

func (t *Toyota) Sum() float64 {
	t.Total = t.Price * t.Discount
	return t.Total
}

func (t *Toyota) SetName(name string) {
	t.Name = name
}

func New(name string) *Toyota {
	car := Toyota{}
	car.Name = name
	car.Discount = 0.8
	car.Price = 4000

	return &car
}
