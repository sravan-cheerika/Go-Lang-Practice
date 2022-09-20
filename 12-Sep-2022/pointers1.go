package main

import "fmt"

func main() {
    //pointers 
    a:=10
    var p = &a
    fmt.Println("value of a :", a)
    fmt.Println("address of a :", &a)
    fmt.Println("value of p :", p)
    fmt.Println("address of p :", *p)

    *p=20

    ptr := new(int) 
    *ptr=100 
    fmt.Println("address of ptr :", *ptr) 
}
