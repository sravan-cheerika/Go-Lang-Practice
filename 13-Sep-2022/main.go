package main

import "fmt"

type Person struct{
    Name string
    Age int
}

func (p Person) GetNameAndAge() Person {
    return Person{p.Name,p.Age}
}

func main(){
    fmt.Println("test : ")
    p1:=Person{"sravan",25}
    fmt.Println("Person name :", p1.GetNameAndAge())
    
}