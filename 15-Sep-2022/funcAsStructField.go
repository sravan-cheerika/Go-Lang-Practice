package main

import "fmt"
//import "reflect"
type Gross func(float32, float32, ) float32
type Average func(int) int

type Student struct{
    Name string
    Age int
    Eng1_sem float32
    Eng2_sem float32
    GrossTotal Gross
    AverageSem Average
}

func Subject(i int,avg Average) int {
    fmt.Println("test")
    return 20;
}
func main() {
    
    Subject(20,s1.AverageSem(20)) 

    s1:=Student{
        Name: "sravan",
        Age: 21,
        Eng1_sem: 99.9,
        Eng2_sem: 99.9,
        GrossTotal:func(f float32, s float32) float32{
            return f+s
        },

    }

    s1:=Student{
        Name: "sravan",
        Age: 21,
        Eng1_sem: 99.9,
        Eng2_sem: 99.9,
        GrossTotal:func(f float32, s float32) float32{
            return f+s
        },
        AverageSem:func(i int){
            return 20
        }

    }


    fmt.Println(s1.Name)
    fmt.Println(s1.GrossTotal(s1.Eng1_sem,s1.Eng2_sem))

}
