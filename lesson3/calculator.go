package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, **, ! - только первое число): ")
	fmt.Scanln(&op)

	isDividerValid(b, op)
	isFactorialValid(a, op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "**":
		res = math.Pow(a, b)
	case "!":
		res = float64(factorial(a))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func isDividerValid(b float64, op string) {
	if b == 0 && op == "/" {
		fmt.Println("На ноль делить нельзя.")
		os.Exit(1)
	}
}

func isFactorialValid(a float64, op string) {
	if a < 0 && op == "!" {
		fmt.Println("Введите первое число больше или равное нулю.")
		os.Exit(1)
	}
}

func factorial(a float64) int {
	num := int(a)
	result := 1

	if a == 0 {
		return result
	} else {
		for i := 1; i <= num; i++ {
			result = result * i
		}
	}

	return result
}
