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
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	// send channel
	go func(ch chan<- int) {
		ch <- 100
		ch <- 101
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
