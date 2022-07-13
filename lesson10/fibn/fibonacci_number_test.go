package fibn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = map[int]int{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
}

func TestFibonacci(t *testing.T) {
	for i := 0; i < len(testCases); i++ {
		assert.Equal(t, testCases[i], Fibonacci(i))
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)

	}
}

func ExampleFibonacci() {
	f := Fibonacci(5)
	fmt.Println(f)

}

func TestFibonacciOpt(t *testing.T) {
	for i := 0; i < len(testCases); i++ {
		assert.Equal(t, testCases[i], FibonacciOpt(i))
	}
}

func BenchmarkFibonacciOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciOpt(10)
	}
}

func ExampleFibonacciOpt() {
	f := FibonacciOpt(5)
	fmt.Println(f)
}
