package main
import "fmt"
import "sort"
//import "strings"

func main(){
    
    people:=[]struct{
        Name string
        Age int
    } {
        {"sravan", 21},
        {"kiran", 34},
        {"charan", 50},
    }

    lessByName:= func(i, j int) bool{
        return people[i].Name < people[j].Name
    }
    lessByAge:=func(i, j int) bool{
        return people[i].Age < people[j].Age
    }
    sort.Slice(people, lessByName )
    fmt.Println("By people", people)

    sort.Slice(people, lessByAge)
    fmt.Println("By age", people)



}