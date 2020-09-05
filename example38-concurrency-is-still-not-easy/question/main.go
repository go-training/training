package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const concurrencyProcesses = 10 // limit the maximum number of concurrent reading process tasks
	const jobCount = 100

	var wg sync.WaitGroup
	wg.Add(jobCount)
	found := make(chan int)
	limitCh := make(chan struct{}, concurrencyProcesses)

	for i := 0; i < jobCount; i++ {
		limitCh <- struct{}{}
		go func(val int) {
			defer func() {
				wg.Done()
				<-limitCh
			}()
			waitTime := rand.Int31n(1000)
			fmt.Println("job:", val, "wait time:", waitTime, "millisecond")
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
			found <- val
		}(i)
	}
	go func() {
		wg.Wait()
		close(found)
	}()
	var results []int
	for p := range found {
		fmt.Println("Finished job:", p)
		results = append(results, p)
	}

	fmt.Println("result:", results)
}
