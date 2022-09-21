package main

import (
	"context"
	"fmt"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		fmt.Println("hi : ")
		dst := make(chan int)
		n := 1
		go func() {
			for{
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				fmt.Println("sending : ")
				n++
			}
		}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println("Receive : ", n)
		if n == 8 {
			break
		}
	}
}
