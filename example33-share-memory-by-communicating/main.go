package main

import (
	"fmt"
	"sync"
)

func addByShareMemory(n int) []int {
	var ints []int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()
		}(i)
	}

	wg.Wait()

	return ints
}

// WriteOnly serves the purpose of demonstrating
// a method that writes to a write-only channel.
func WriteOnly(channel chan<- int, order int) {
	channel <- order
}

func addByShareCommunicate(n int) []int {
	var ints []int
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go WriteOnly(channel, i)
	}

	for i := range channel {
		ints = append(ints, i)

		if len(ints) == n {
			break
		}
	}

	close(channel)

	return ints
}

func main() {
	foo := addByShareMemory(10)
	fmt.Println(len(foo))
	fmt.Println(foo)

	foo = addByShareCommunicate(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}
