package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Let's Go")
	}()
	// fmt.Println("exit!!!!")
}
