package main
import "fmt"
//import "sort"
import "strings"

func Filter(data[] string,fil func(string ) bool, ) ([] string){
    found:=make([] string, 0)
    for _,element:=range data{
        if( fil(element)){
            found=append(found,element)
        }
    }
        return found
}

func main(){
    word:=[]string{"India","Europe","US"}
    sp:="I"
    output:= Filter(word, func(s string) bool{
        return strings.HasPrefix(s, sp)
    })
    fmt.Println(output)

    sp:="In"
    output:= Filter(word, func(s string) bool{
        return strings.contains()
    })

}