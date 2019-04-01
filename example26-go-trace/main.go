package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func Sum(numbers ...int) int {

	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SumParallel(numbers ...int) int {
	mid := len(numbers) / 2

	ch := make(chan int)
	go func() { ch <- Sum(numbers[:mid]...) }()
	go func() { ch <- Sum(numbers[mid:]...) }()

	total := <-ch + <-ch
	return total
}

var foo = []int{4, 5, 6, 1, 3, 2, 8, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	fmt.Println(runtime.NumCPU())
	// fmt.Println(Sum(foo...))
	fmt.Println(SumParallel(foo...))
}
