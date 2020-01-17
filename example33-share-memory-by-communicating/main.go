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

func main() {
	foo := addByShareMemory(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}
