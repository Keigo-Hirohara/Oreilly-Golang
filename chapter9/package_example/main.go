package main

import (
	"chapter9/package_example/math"

	"fmt"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)

}
