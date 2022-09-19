package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make(chan int, 2) // Query : to ask
	go func() {
		fmt.Println("reading 1 :", <-numbers)
		fmt.Println("reading 2 :", <-numbers)
		fmt.Println("reading 3 :", <-numbers)
	}()

	time.Sleep(1 * time.Second)
	numbers <- 1
	time.Sleep(1 * time.Second)
	numbers <- 2
	time.Sleep(1 * time.Second)
	numbers <- 3
}
