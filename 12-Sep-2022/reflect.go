package main

import(

    "fmt"
    "reflect"

)

func main(){

    a:=[]int {1,2,3}

    fmt.Println(reflect.ValueOf(a).Kind())
    fmt.Println(fmt.Printf("%s",reflect.ValueOf(a).Kind()))
    fmt.Printf("%s",reflect.ValueOf(a).Kind())

}