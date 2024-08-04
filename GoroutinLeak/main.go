package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			}
		}
	}()
	time.Sleep(1 * time.Second)
}
