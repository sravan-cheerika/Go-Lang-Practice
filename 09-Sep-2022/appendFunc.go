package main

import "fmt"

var fns[] func(i int)

func main(){
fns = append(fns,one)
fns = append(fns,two)

for _,fn:=range fns{
    fn()
}
}

func one(){
    fmt.Println("one")
}
func two(){
    fmt.Println("two")
}