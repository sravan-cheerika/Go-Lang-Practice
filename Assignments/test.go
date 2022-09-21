package main

import (
    "fmt"
    "time"
)

func main() {

    c, b := make(chan int), make(chan rune)
    go func() {
        for i := 0; i < 6; i++ {
            c <- i
			time.Sleep(1* time.Millisecond)
        }

    }()

    go func() {

        for j := 'a'; j <= 'e'; j++ {
            b <- j
			time.Sleep(1* time.Millisecond)
        }

    }()

	for i := 0; i < 10; i++ {
        select {
        case msg1 := <-c:
            fmt.Println("received 1", string(msg1))
        case msg2 := <-b:
            fmt.Println("received 2", msg2)
        }
    }
 //anynoums function call
}