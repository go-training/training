package main

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// how to fix the timeout issue?
	for {
		select {
		case val := <-output:
			if ctx.Err() != nil {
				println("reached timeout, but still have data to process")
				return
			}
			// simulate slow consumer
			time.Sleep(500 * time.Millisecond)
			println("output:", val)
		// how to fix the timeout issue?
		case <-ctx.Done():
			println("timeout")
			return
		}
	}
}
