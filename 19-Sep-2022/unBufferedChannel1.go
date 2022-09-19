package main

import (
	"fmt"
	"time"
)

func access(ch chan int) { // unBuffered since no value mentioned..
	time.Sleep(time.Second)
	fmt.Println("start accessing channel\n")
// 2 ways of for loops in Go -> 1. condition based & 2. using range of slice
	for i := range ch {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int, 2)
	defer close(ch) // defer used to execute the func at the last of main()
	// defer close1() //

	go access(ch)

	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Filled")
	}

	time.Sleep(1 * time.Second)
}
