package main

import "fmt"

func main() {
    a:= make([] int,7,9) // 
/*    
    a=append(a,1)
    fmt.Println("a is : ", a)
    a=append(a,2)
    fmt.Println("lates a : ", a)
*/
    for i:=1;i<=10;i++{
        a=append(a,i)
        fmt.Println("latest a : ", a, len(a), cap(a))
    }
}
