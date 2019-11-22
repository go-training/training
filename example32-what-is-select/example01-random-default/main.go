package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("exit")
	}
}
