package main

import "fmt"

type Employee struct{
    FirstName, LastName string
    Age int
}
func main() {
    //pointers 
    a:=10
    c:=&a

    fmt.Println("c is : ",c)
    fmt.Println("c is : ",*c)

    e:=&Employee{"sravan","reddy",23}

    fmt.Println("e.FN", e.FirstName)

/*    var p = &a

    var pp = &p

    fmt.Println("address of *p :", *p, "value of *pp", *pp )
    fmt.Println("value of **pp", **pp )
*/
}
