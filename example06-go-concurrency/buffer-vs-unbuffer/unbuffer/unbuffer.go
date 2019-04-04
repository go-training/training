package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		c <- true
	}()
	<-c
}
