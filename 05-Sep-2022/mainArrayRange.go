package main

import "fmt"

func main() {
    names:=[3] string {"sravan","reddy","cheerika"}
    for index,val := range(names){
        fmt.Println(index," : ",val)
    }
}
