package main

import "fmt"

func main() {
	fmt.Println(div(10, 5))
	fmt.Println(div(6, 0))
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
