package main

import (
	"time"
)

func main() {
	output := make(chan int, 30)

	go func() {
		for i := 0; i < 30; i++ {
			output <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	timeout := time.After(1 * time.Second)
	// how to fix the timeout issue?
	for {
		select {
		case val := <-output:
			select {
			case <-timeout:
				println("reached timeout, but still have data to process")
				return
			default:
			}
			// simulate slow consumer
			time.Sleep(500 * time.Millisecond)
			println("output:", val)
		// how to fix the timeout issue?
		case <-timeout:
			println("timeout")
			return
		}
	}
}
