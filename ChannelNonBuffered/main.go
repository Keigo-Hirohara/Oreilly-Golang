package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	var v2 int
	select {
	case ch2 <- v: // サブゴルーチンによってch2が読み込まれたら、vの値をch2に書き込み
	case v2 = <-ch1: // サブゴルーチンによってch1に値が書き込まれたら、v2に読み込む
	}
	fmt.Println(v, v2)
}
