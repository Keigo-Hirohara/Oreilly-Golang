package main

import (
	"errors"
	"fmt"
	"os"
)

func calcRemaindAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}

	return numerator / denominator, numerator % denominator, nil
}

func main() {
	numerator, denominator := 20, 3
	remainder, mod, error := calcRemaindAndMod(numerator, denominator)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Printf("%d÷%d: 商: %d, 余り: %d\n", numerator, denominator, remainder, mod)
}
