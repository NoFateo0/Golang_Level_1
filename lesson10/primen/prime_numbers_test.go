package primen

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = map[int][]int{
	0:  nil,
	1:  nil,
	2:  {2},
	3:  {2, 3},
	4:  {2, 3},
	5:  {2, 3, 5},
	6:  {2, 3, 5},
	7:  {2, 3, 5, 7},
	8:  {2, 3, 5, 7},
	9:  {2, 3, 5, 7},
	10: {2, 3, 5, 7},
}

func TestPrimeN(t *testing.T) {
	for i := 0; i <= 10; i++ {
		t.Run(fmt.Sprintf("n = %d", i), func(t *testing.T) {
			if !reflect.DeepEqual(PrimeN(i), testCases[i]) {
				t.Errorf("Неверный результат для n = %d", i)
			}
		})
	}
}

func ExamplePrimeN() {
	f := PrimeN(10)
	fmt.Println(f)
}
