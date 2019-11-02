package main

import (
	"fmt"
	"time"
)

func worker(jobChan <-chan int) {
	for job := range jobChan {
		fmt.Println("current job:", job)
		time.Sleep(3 * time.Second)
		fmt.Println("finished job:", job)
	}
}

func main() {
	// make a channel with a capacity of 1.
	jobChan := make(chan int, 1)

	// start the worker
	go worker(jobChan)

	// enqueue a job
	fmt.Println("enqueue the job 1")
	jobChan <- 1
	fmt.Println("enqueue the job 2")
	jobChan <- 2
	fmt.Println("enqueue the job 3")
	go func() {
		jobChan <- 3
	}()

	fmt.Println("waiting the jobs")
	time.Sleep(10 * time.Second)
}
