package main

import "fmt"

// The Time Complexity of the Bubble Sort is O(n^2) since it takes two nested loops to check the adjacent element.
// For example, let’s take the following unsorted array −
// 22 15 11 45 13
// Bubble Sort Algorithm first traverses the whole array and then in another loop checks if the adjacent elements are in order or not.
// Thus, after sorting the elements will be,
// 11 13 15 22 45

// conver to generics type to support int and float64 types
func bubbleSort(array []int) []int {
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
	array := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(bubbleSort(array))
}
