package main

import "fmt"
var enter string

func main() {
    fmt.Println("Hello world")
    fmt.Scan(&enter)
    if(enter=="hey"){
        fmt.Println("success")
    } else{
        fmt.Println("failure")
    }
}
