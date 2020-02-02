package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		ch <- 1
		ch <- 2
		// close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}

	// go func() {
	// 	for n := range ch {
	// 		fmt.Println(n)
	// 	}
	// }()

	time.Sleep(1 * time.Second)
}
