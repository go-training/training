package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

var (
	m1 = map[int]int{1: 2, 2: 4, 4: 8, 8: 16}
	m2 = map[int]string{1: "2", 2: "4", 4: "8", 8: "16"}
)

func main() {
	fmt.Println(maps.Keys(m1))
	fmt.Println(maps.Keys(m2))

	fmt.Println(maps.Values(m1))
	fmt.Println(maps.Values(m2))

	fmt.Println(maps.Equal(m1, map[int]int{1: 2, 2: 4, 4: 8, 8: 16}))

	maps.Clear(m1)
	fmt.Println(m1)
	m3 := maps.Clone(m2)
	fmt.Println(m3)
}
