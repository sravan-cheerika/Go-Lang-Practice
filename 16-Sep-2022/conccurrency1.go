package main

import ( 
    "fmt"
    "time"
)
func foo() {
    for i:=0; i<15; i++{
        fmt.Println("foo : ", i)
        //time.Sleep(1*time.Millisecond)
    }
}
func bar() {
    for i:=0; i<10; i++{
        fmt.Println("bar : ", i)
        //time.Sleep(1*time.Millisecond)
    }
}
func main() {
    go foo() // signifies goRoutines for multi threading concept in GoLang. Execute functions simultaneously 
    go bar()

    time.Sleep(1*time.Second)
    //fmt.Println("hi")

}
