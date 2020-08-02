package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan bool)
	go func() {
		for {
			select {
			case <-exit:
				fmt.Println("Exit")
				return
			case <-time.After(2 * time.Second):
				fmt.Println("Monitoring")
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Notify Exit")
	exit <- true //keep main goroutine alive
	time.Sleep(5 * time.Second)
}
