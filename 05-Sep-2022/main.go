package main

import "fmt"

type Swap(&x, &y) func{
    &x,y = y,x    
}

func main() {
    //pointers
    a:=10
    b:=20

    c:=&a
    d:=&b

    fmt.Println(" before swap :", a, b)
    Swap(a,b)

    fmt.Println(" after swap :",a,b)

    fmt.Println("c is : ",c)
    fmt.Println("*c is : ",*c)

    e:=&Employee{"sravan","reddy",23}

    fmt.Println("e.FN", e.FirstName)

/*    var p = &a

    var pp = &p

    fmt.Println("address of *p :", *p, "value of *pp", *pp )
    fmt.Println("value of **pp", **pp )
*/
}
