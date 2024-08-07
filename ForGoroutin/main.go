package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch := make(chan int)
	for _, v := range a {
		go func(input int) {
			ch <- input * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch, " ")
	}
}
