package main

import "fmt"

func fibonacci(n int, fibonacciMap map[int]int) int {
	if fibonacciMap == nil {
		fibonacciMap = make(map[int]int, n)
		fibonacciMap[0] = 0
		fibonacciMap[1] = 1
	}

	result, isExist := fibonacciMap[n]

	if !isExist {
		result = fibonacci(n-2, fibonacciMap) + fibonacci(n-1, fibonacciMap)
		fibonacciMap[n] = result
	}

	return result
}

func main() {
	var n int
	fmt.Print("Введите номер числа: ")
	fmt.Scan(&n)
	fmt.Printf("Значение вашего числа: %d", fibonacci(n, nil))
}
