package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	i := 0
	wg.Add(3) //task count wait to do
	go func() {
		defer wg.Done() // finish task1
		fmt.Println("goroutine 1 done")
		i++
	}()
	go func() {
		defer wg.Done() // finish task2
		fmt.Println("goroutine 2 done")
		i++
	}()
	go func() {
		defer wg.Done() // finish task3
		fmt.Println("goroutine 3 done")
		i++
	}()
	wg.Wait() // wait for tasks to be done
	fmt.Println("all goroutine done")
	fmt.Println(i)
}
