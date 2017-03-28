package lexus

type Lexus struct {
	name     string
	price    float64
	discount float64
	color    string
}

func (l Lexus) Name() string {
	return l.name
}

func (l Lexus) Price() float64 {
	return l.price * l.discount
}

func (l Lexus) Discount() float64 {
	return l.discount
}

func (l Lexus) Color() string {
	return l.color
}

// NewCar contructure
func NewCar(name string, price float64, discount float64, color string) *Lexus {
	return &Lexus{
		name:     name,
		price:    price,
		discount: discount,
		color:    color,
	}
}
