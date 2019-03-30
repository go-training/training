package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	// receive channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	// send channel
	go func(ch chan<- int) {
		ch <- 100
		ch <- 101
		wg.Done()
	}(ch)
	wg.Wait()
}
