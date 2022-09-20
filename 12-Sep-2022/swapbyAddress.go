package main

import "fmt"

func swap(x int, y int){
    x,y = y,x
}

func swapByAddress(x *int, y *int){
    *x,*y = *y,*x
}

func main() {
    //pointers 
    a,b:=10,20
    fmt.Println(a,b)
    swap(a,b)
    fmt.Println("swap without address "a,b)
    swapByAddress(&a, &b)
    fmt.Println("swap with address "a,b)

    fmt.Println("after swap() : ", a,b)
    
}
