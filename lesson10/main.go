package main

import (
	"fmt"
	"lesson10/fibn"
	"lesson10/primen"
	inserts "lesson10/sorts"
)

func main() {
	var num int
	var slice = []int{1, 2, 5, -3, -10, 3, -7, 11, 6, 0}
	fmt.Print("Введите ваше число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Некоректный ввод. Введите число заново.")
	} else {
		fmt.Println("Простые числа до введенного числа: ", primen.PrimeN(num))
		fmt.Printf("Значение введенного числа Фибонначи: %d\n", fibn.FibonacciOpt(num))

	}
	fmt.Println("Слайс для сортировки: ", slice)
	fmt.Println("Сортировка вставкой: ", inserts.InsertionSort([]int{1, 2, 5, -3, -10, 3, -7, 11, 6, 0}))
	fmt.Println("Сортировка пузырьком: ", inserts.BubbleSort([]int{1, 2, 5, -3, -10, 3, -7, 11, 6, 0}))
}
