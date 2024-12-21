package main

import (
	"time"
)

func main() {
	output := make(chan int, 10)

	go func() {
		for i := 0; i < 30; i++ {
			output <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// how to fix the timeout issue?
	for {
		select {
		case val := <-output:
			// simulate slow consumer
			time.Sleep(500 * time.Millisecond)
			println("output:", val)
		// how to fix the timeout issue?
		case <-time.After(1 * time.Second):
			println("timeout")
			return
		}
	}
}
