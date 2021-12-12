package main

import (
	"fmt"
)

func main() {
	var sum int
	fmt.Scanln(&sum)
	if sum > 1000 {
		fmt.Println("Apple")
	} else if sum >= 500 && sum <= 1000 {
		fmt.Println("Samsung")
	} else if sum < 500 {
		fmt.Println("Nokia с фонариком")
	}

}
