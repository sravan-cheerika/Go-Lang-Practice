package main
import "fmt"
import "sort"

type Student struct {
    
}

func main(){
    word := [] string {"cat","aab","abc","apple","bat"}
    sort.Strings(word)
    fmt.Println(word)

    sort.Sort( sort.Reverse( sort.StringSlice(word)))
    fmt.Println(word)
   /* s := []int{5,2,9,1,0,3}
    sort.Ints(s)
    fmt.Println(s)
    */
}