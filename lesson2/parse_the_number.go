package main

import "fmt"

func main() {
	var number int

	fmt.Println("Введите число от 0 до 999: ")
	fmt.Scanln(&number)

	if number >= 1000 || number < 0 {
		fmt.Println("Вы ввели неправильное число, попробуйте снова.")
	} else {
		hundreds := number / 100
		tens := number % 100 / 10
		ones := number % 10
		fmt.Printf("В числе %d содержится %d сотни(ен) %d десятка(ов) %d единиц(ы).", number, hundreds, tens, ones)
	}
	
}
