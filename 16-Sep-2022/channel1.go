package main
import (
    "fmt"
    //"sync"
    "time"
    )

func main() {
    c := make(chan int) // c <- writing(c on left side) && -> c reading (c on right side)

    go func() {
        for i:=0; i<=10; i++{
            c <- i // sending value through channel 
        }
    }()
    go func() {
        for {
            fmt.Println(<-c)
        }
    }()
    go func() {
        for i:=1;i<=2;i++ {
            fmt.Println("hi : ")
        }
    }()

    time.Sleep(time.Second)
}