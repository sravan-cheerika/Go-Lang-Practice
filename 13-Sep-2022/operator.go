package main

import "fmt"
import "strconv"


func main() {
    var a int = 60
    var b int = 50
    var c int = 0

    //a,b,c := 10,20,30 

    c = a | b
    c = a ^ b

    a =1
    b = a<<1
    //show(a,b,c)
    c=a>>1 

    fmt.Println("Hello world")
}
