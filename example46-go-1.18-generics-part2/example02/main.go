package main

import "fmt"

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

type car struct {
	price int
}

func (c car) String() string {
	return fmt.Sprintf("%d", c.price)
}

type plusser interface {
	Plus(string) string
}

func cconcatTo[S stringer, P plusser](s []S, p []P) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = p[i].Plus(v.String())
	}
	return r
}

type foo struct {
	name string
}

func (c foo) String() string {
	return c.name
}

type plus map[string]string

func (p plus) Plus(k string) string {
	if v, ok := p[k]; ok {
		return v
	}
	return ""
}

func main() {
	// exmaple01
	val := concat([]stringer{
		car{price: 1},
		car{price: 2},
	})

	fmt.Println(val)

	// example02
	p := plus{
		"a": "100",
		"b": "200",
	}

	fmt.Println(cconcatTo([]stringer{
		foo{name: "a"},
		foo{name: "b"},
	}, []plusser{p, p}))
}
