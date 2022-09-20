package main

import (
    "fmt"
    "math"
)

type circle struct{
    radius float64
}

type Triangle struct{
    Base float64
    Height float64
}

type shape interface{
    area() float64
}

func (c circle) area() float64 {
//func  area() (c circle) float64 { For *passing struct we need it on the left side of func name : 
    return math.Pi * c.radius * c.radius
}

func info(s shape) {
    fmt.Println("area", s.area())
}

func (t Triangle) area() float64{
    return t.Base * t.Height
}
func main() {
    c:=circle{5}
    //fmt.Println(": ")
    info(c)
    tt:=Triangle{10,5}
    info(tt)

}
