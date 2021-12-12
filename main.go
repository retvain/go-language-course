package main

import (
	"fmt"
)

func mars_age(a int) float64 {
	var earthDays, marsAge int
	earthDays = a * 365
	marsAge = earthDays / 687

	return float64(marsAge)
}

func main() {
	fmt.Println(mars_age(1000))
}
