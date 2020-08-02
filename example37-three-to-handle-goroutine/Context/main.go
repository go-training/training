package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, name string) {
	go bar(ctx, name) // A calls B
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "A Exit")
			return
		case <-time.After(1 * time.Second):
			fmt.Println(name, "A do something")
		}
	}
}

func bar(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "B Exit")
			return
		case <-time.After(2 * time.Second):
			fmt.Println(name, "B do something")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx, "FooBar")
	fmt.Println("client release connection, need to notify A, B exit")
	time.Sleep(5 * time.Second)
	cancel() //mock client exit, and pass the signal, ctx.Done() gets the signal  time.Sleep(3 * time.Second)
	time.Sleep(3 * time.Second)
}
