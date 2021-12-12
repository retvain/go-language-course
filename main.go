package main

import "fmt"

func main() {
	//Ваш робот должен принимать на вход 3 целых числа (каждый с новой строки) в диапазоне от 0 до 10 (включительно) и выдавать соответствующий текст на русском языке.
	var num1 int
	for i := 1; i <= 3; i++ {
		fmt.Scanln(&num1)
		switch num1 {
		case 0:
			fmt.Println("Ноль")
		case 1:
			fmt.Println("Один")
		case 2:
			fmt.Println("Два")
		case 3:
			fmt.Println("Три")
		case 4:
			fmt.Println("Четыре")
		case 5:
			fmt.Println("Пять")
		case 6:
			fmt.Println("Шесть")
		case 7:
			fmt.Println("Семь")
		case 8:
			fmt.Println("Восемь")
		case 9:
			fmt.Println("Девять")
		case 10:
			fmt.Println("Десять")
		}
	}
}
