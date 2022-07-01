package main

import "fmt"

func main() {
	array := []int{1, 2, 5, -3, -10, 3, -7, 11, 6, 0}
	fmt.Println(array)
	insertionSort(array)
	fmt.Println(array)
}

func insertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		var t int
		for j := i; j > 0 && (array[j] < array[j-1]); j-- {
			t = array[j-1]
			array[j-1] = array[j]
			array[j] = t
		}
	}
}
