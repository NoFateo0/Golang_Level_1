package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Введите длину: ")
	fmt.Scanln(&a)

	fmt.Print("Введите ширину: ")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника равна:", a*b)

}
