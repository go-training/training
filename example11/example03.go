package main

import "fmt"

func printValue(v interface{}) {
	switch v := v.(type) {
	case string:
		fmt.Printf("%v is a string\n", v)
	case int:
		fmt.Printf("%v is a int\n", v)
	default:
		fmt.Printf("The type of v is unknown\n")
	}
}

func printValue2(v ...interface{}) {
	for _, v := range v {
		switch v := v.(type) {
		case string:
			fmt.Printf("%v is a string\n", v)
		case int:
			fmt.Printf("%v is a int\n", v)
		default:
			fmt.Printf("The type of v is unknown\n")
		}
	}
}

func main() {
	printValue(10)
	printValue("100")
	printValue(123.01)
	printValue2(300, "400", 500.01)
}
