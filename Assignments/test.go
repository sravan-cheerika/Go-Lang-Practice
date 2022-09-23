package main

import (
    "fmt"
    "time"
)

func main() {

    c, b := make(chan int), make(chan rune)
    go func() {
        for i := 1; i < 35; i++ {
            c <- i
			time.Sleep(1* time.Millisecond)
        }

    }()

    go func() {

        for j := 'A'; j <= 'Z'; j++ {
            b <- j
			time.Sleep(1* time.Millisecond)
        }

    }()

	go func() {

        for {
            fmt.Print(<-c)
            fmt.Print(string(<-b))
        }

    }() //anynoums function call

    time.Sleep(time.Second)
 //anynoums function call
}