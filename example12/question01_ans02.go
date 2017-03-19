package main

import (
	"fmt"
)

// create new channel
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Let's Go")
		c <- true
	}()

	<-c
}
