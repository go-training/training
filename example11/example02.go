package main

import "fmt"

func printValue(v interface{}) {
	if v, ok := v.(string); ok {
		fmt.Printf("The value of v is: %#v", v)
	} else {
		fmt.Println("Oops, it is not a string!")
	}
}

func main() {
	v := "10"
	printValue(v)
}
