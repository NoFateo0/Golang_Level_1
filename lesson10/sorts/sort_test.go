package inserts

import (
	"fmt"
	"reflect"
	"testing"
)

var slice = []int{1, 2, 5, -3, -10, 3, -7, 11, 6, 0}

var testCases = []struct {
	name   string
	input  []int
	result []int
}{
	{
		name:   "Пустой слайс",
		input:  []int{},
		result: []int{},
	},
	{
		name:   "Сортированный слайс",
		input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		result: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name:   "Слайс отрицательных чисел",
		input:  []int{-10, -20, -1, -3, -5, -33, -100, -99, -50, -1000},
		result: []int{-1000, -100, -99, -50, -33, -20, -10, -5, -3, -1},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, cse := range testCases {
		t.Run(cse.name, func(t *testing.T) {
			if !reflect.DeepEqual(InsertionSort(cse.input), cse.result) {
				t.Errorf("Неверный результат для: \"%s\"", cse.name)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(slice)
	}
}

func ExampleInsertionSort() {
	unsortedSlice := []int{10, 2, 6, 1}
	s := InsertionSort(unsortedSlice)
	fmt.Println(s)
}

func TestBubbleSort(t *testing.T) {
	for _, cse := range testCases {
		t.Run(cse.name, func(t *testing.T) {
			if !reflect.DeepEqual(BubbleSort(cse.input), cse.result) {
				t.Errorf("Неверный результат для: \"%s\"", cse.name)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(slice)
	}
}

func ExampleBubbleSort() {
	unsortedSlice := []int{10, 2, 6, 1}
	s := BubbleSort(unsortedSlice)
	fmt.Println(s)
}
