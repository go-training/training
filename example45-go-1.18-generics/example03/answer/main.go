package main

import "fmt"

// The Time Complexity of the Bubble Sort is O(n^2) since it takes two nested loops to check the adjacent element.
// For example, let’s take the following unsorted array −
// 22 15 11 45 13
// Bubble Sort Algorithm first traverses the whole array and then in another loop checks if the adjacent elements are in order or not.
// Thus, after sorting the elements will be,
// 11 13 15 22 45

// conver to generics type to support int and float64 types

type Number interface {
	int | int32 | int64 | float32 | float64
}

func bubbleSort[n Number](array []n) []n {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	n1 := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(bubbleSort(n1))
	n2 := []float64{11.1, 14.2, 3.3, 8.4, 18.5, 17.6, 43.7}
	fmt.Println(bubbleSort(n2))
}
