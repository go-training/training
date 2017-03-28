package toyota

type Toyota struct {
	name     string
	price    float64
	discount float64
	color    string
}

func (t Toyota) Name() string {
	return t.name
}

func (t Toyota) Price() float64 {
	return t.price * t.discount
}

func (t Toyota) Discount() float64 {
	return t.discount
}

func (t Toyota) Color() string {
	return t.color
}

// NewCar contructure
func NewCar(name string, price float64, discount float64, color string) *Toyota {
	return &Toyota{
		name:     name,
		price:    price,
		discount: discount,
		color:    color,
	}
}
