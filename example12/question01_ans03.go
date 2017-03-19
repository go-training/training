package main

import (
	"fmt"
)

// Close channel
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Let's Go")
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
