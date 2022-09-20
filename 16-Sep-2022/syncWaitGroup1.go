package main

import ( 
    "fmt"
    "time"
    "sync"
)

var wg sync.WaitGroup 
func foo() {
    for i:=0; i<5; i++{
        fmt.Println("Foo : ", i)
        time.Sleep(2 * time.Millisecond)
    }
    wg.Done()
}

func bar() {
    for i:=0; i<5; i++{
        fmt.Println("Bar : ", i)
        time.Sleep(1 * time.Millisecond)
    }
    wg.Done()
}

func main() {
    wg.Add(2) // no of func to execute 
    go foo()
    go bar()
    fmt.Println("hi")
    wg.Wait()
}
