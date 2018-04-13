package main

import "fmt"

// pass slice as function argument
func addValue(foo []string) {
	foo = append(foo, "c")
	fmt.Println("modify foo", foo)
}

func main() {
	foo := []string{"a", "b"}
	fmt.Println("before foo:", foo)
	addValue(foo)
	fmt.Println("after foo:", foo)
}
