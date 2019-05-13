package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(val int, wg *sync.WaitGroup) {
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			fmt.Println("finished job id:", val)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
