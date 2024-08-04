package main

import "fmt"

func main() {
	const maxConc = 10
	ch := make(chan int, maxConc)

	for i := 0; i < maxConc; i++ {
		ch <- i
	}
	close(ch)

	result := processChannel(ch, maxConc)

	fmt.Println(result)
}

func processChannel(ch chan int, maxConc int) []int {
	results := make(chan int, maxConc)
	for i := 0; i < maxConc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	fmt.Println("ゴルーチン起動完了")

	var out []int

	for i := 0; i < maxConc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(input int) int {
	return input * 100
}
