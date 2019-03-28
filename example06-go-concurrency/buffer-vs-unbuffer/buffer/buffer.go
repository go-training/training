package main

import "fmt"

// how to show the "GO GO GO"
func main() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("GO GO GO")
		<-c
	}()
	c <- true
}
