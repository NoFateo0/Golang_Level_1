package fibn

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}
func FibonacciOpt(n int) int {
	var fibonacciMap = map[int]int{
		0: 0,
		1: 1,
	}
	if val, ok := fibonacciMap[n]; ok {
		return val
	}

	fibonacciMap[n] = FibonacciOpt(n-2) + FibonacciOpt(n-1)

	return fibonacciMap[n]
}
