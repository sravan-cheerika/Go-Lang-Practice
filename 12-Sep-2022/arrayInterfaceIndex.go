package main

import(
    "fmt"
)
type Element interface{}
type List [] Element

type Person struct{
    name string
    age int
} 

func main(){
    list:= make(List,4)
    list[0] = 1
    list[1] = "sravan"

    for index, element := range list{
        if value, ok:= element.(int); ok {
            fmt.Println("inside int: ",index, value, element)
        } else if value,ok := element.(string); ok{
            fmt.Println("inside string: ",index, value, element)

        }
    }
}