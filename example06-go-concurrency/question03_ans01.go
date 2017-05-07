package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Let's Go"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Let's GoGoGo"
	time.Sleep(1 * time.Second)
}
