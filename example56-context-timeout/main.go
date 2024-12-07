package main

import (
	"time"
)

func main() {
	output := make(chan int, 1)

	go func() {
		for i := 0; i < 30; i++ {
			output <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case val := <-output:
			println("output:", val)
		// how to fix the timeout issue?
		case <-time.After(1 * time.Second):
			println("timeout")
			return
		}
	}
}
