package main

import "fmt"

func main() {
	//На вход подается целое число. Вам необходимо вывести произведение всех чисел от 1 до данного числа (включительно).
	//Например, на вход подается число 5, вам нужно найти - 1*2*3*4*5 = 120
	var num, sum int
	fmt.Scanln(&num)
	sum = 1
	for i := 1; i < num; i++ {
		sum = sum * (i + 1)
	}
	fmt.Println(sum)
}
