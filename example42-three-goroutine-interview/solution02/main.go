package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomTime() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func main() {
	ch01 := make(chan struct{})
	ch02 := make(chan struct{})
	ch03 := make(chan struct{})
	count := 10
	wg := &sync.WaitGroup{}
	wg.Add(count)

	go func() {
		ch01 <- struct{}{}
	}()
	go func(in, out chan struct{}) {
		for i := 0; i < count; i++ {
			<-in
			time.Sleep(time.Duration(randomTime()) * time.Millisecond)
			fmt.Println("cat")
			out <- struct{}{}
		}
	}(ch01, ch02)

	go func(in, out chan struct{}) {
		for i := 0; i < count; i++ {
			<-in
			time.Sleep(time.Duration(randomTime()) * time.Millisecond)
			fmt.Println("dog")
			out <- struct{}{}
		}
	}(ch02, ch03)

	go func(in, out chan struct{}) {
		for i := 0; i < count; i++ {
			<-in
			time.Sleep(time.Duration(randomTime()) * time.Millisecond)
			fmt.Println("bird")
			if i != (count - 1) {
				out <- struct{}{}
			}
			wg.Done()
		}
	}(ch03, ch01)

	wg.Wait()
}
