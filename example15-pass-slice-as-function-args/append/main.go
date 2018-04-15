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
	bar := foo[:1]
	fmt.Println("bar:", bar)
	s1 := append(bar, "c")
	fmt.Println("foo:", foo)
	fmt.Println("s1:", s1)
	s2 := append(bar, "d")
	fmt.Println("foo:", foo)
	fmt.Println("s2:", s2)
	s3 := append(bar, "e", "f")
	fmt.Println("foo:", foo)
	fmt.Println("s3:", s3)
}
