package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	time.Sleep(5 * time.Second)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
