package main

import (
	"fmt"
	"sync"
)

func cacl(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	t := 0
	for i := 1; i < 1000000; i++ {
		t++
	}

	fmt.Println(i, t)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go cacl(i, &wg)
	}

	wg.Wait()
	fmt.Println("End!!!!")
}
