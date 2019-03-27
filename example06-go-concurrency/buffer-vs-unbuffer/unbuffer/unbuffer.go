package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		<-c
	}()
	c <- true
	time.Sleep(1 * time.Second)
}
