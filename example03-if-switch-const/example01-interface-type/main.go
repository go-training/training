package main

var (
	i interface{}
)

func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is interger", t)
	case string:
		println("i is string", t)
	case float64:
		println("i is float64", t)
	}
}

func main() {
	i = 100
	convert(i)
	i = float64(45.55)
	convert(i)
	i = "foo"
	convert(i)
}
