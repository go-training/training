package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 100
		// copy the i value to channel
		ch <- i
		i = 101
		wg.Done()
	}()
	wg.Wait()
}
