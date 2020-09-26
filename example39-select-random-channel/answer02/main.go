package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			if v, ok := <-ch; ok {
				fmt.Printf("val:%d\n", v)
			}
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 30; i++ {
		ch <- i
		select {
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		default:
		}

		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}
