package main

import (
	"constraints"
	"fmt"
)

// ToString convert any type to string
func ToString(value interface{}) string {
	if v, ok := value.(*string); ok {
		return *v
	}
	return fmt.Sprintf("%v", value)
}

func toString[T constraints.Ordered](value T) string {
	return fmt.Sprintf("%v", value)
}

// ToBool convert any type to boolean
func ToBool(value interface{}) bool {
	switch value := value.(type) {
	case bool:
		return value
	case int:
		if value != 0 {
			return true
		}
		return false
	}
	return false
}

// can't work with type convert
// func toBool[T constraints.Ordered](value T) bool {
// 	switch value := value.(type) {
// 	case bool:
// 		return value
// 	case int:
// 		if value != 0 {
// 			return true
// 		}
// 		return false
// 	}
// 	return false
// }

func main() {
	fmt.Println(ToString("abc"))
	fmt.Println(ToString(1234))
	fmt.Println(ToString(1234.5678))

	fmt.Println(toString("abc"))
	fmt.Println(toString(1234))
	fmt.Println(toString(1234.5678))

	fmt.Println(ToBool(true))
	fmt.Println(ToBool(false))
	fmt.Println(ToBool(1234))
	fmt.Println(ToBool(0))
}
