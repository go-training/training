package main

import (
	"fmt"
	"time"
)

// add time sleep func
func main() {
	go func() {
		fmt.Println("Let's Go")
	}()

	time.Sleep(1 * time.Second)
}
