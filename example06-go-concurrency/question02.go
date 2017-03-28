package main

import (
	"fmt"
)

func cacl(i int, c chan bool) {
	t := 0
	for i := 1; i < 1000000; i++ {
		t++
	}

	fmt.Println(i, t)

	if i == 9 {
		c <- true
	}
}

func main() {
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go cacl(i, c)
	}

	<-c
	fmt.Println("End!!!!")
}
