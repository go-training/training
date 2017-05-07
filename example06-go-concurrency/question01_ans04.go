package main

import (
	"fmt"
	"sync"
)

// add time sleep func
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Let's Go")
	}()
	wg.Wait()
}
