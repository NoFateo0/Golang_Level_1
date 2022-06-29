package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64

	fmt.Println("Введите площадь: ")
	fmt.Scanln(&s)

	d := math.Sqrt(s/math.Pi) * 2
	c := math.Sqrt(s/math.Pi) * 2 * math.Pi
	fmt.Printf("Диметр будет равен: %.5f\n", d)
	fmt.Printf("Длина окружности будет равна: %.5f\n", c)

}
