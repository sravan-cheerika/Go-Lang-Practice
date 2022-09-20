package main
import "fmt"

func main(){
    fmt.Println("hi")
    
    Dict:=map[string] string{}
    Dict["A"]="Mango"
    Dict["B"]="Big Mango"
    Dict["C"]="Apple"

    for key,val:=range Dict{
        fmt.Println(key, " ", val, len(Dict))
    }
    fmt.Println(Dict["A"])
    val,err:=Dict["Z"]
    
    fmt.Println(val," ", err)

}