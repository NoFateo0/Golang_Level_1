package main

import "fmt"

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	for i := 2; i < a; i++ {
		flag := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
			}
		}
		if flag == true {
			continue
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
