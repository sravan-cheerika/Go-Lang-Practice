package main
import (
    "fmt"
//    "time"
    )

func First(c chan<- string){
    for i:=0 ; ;i++ {
        c <- "first"
    }
}

func Second(c chan <- string){
    for i:=0 ; ;i++ {
        c <- "second"
    }
}

/*func PrintFS(c <- chan string){
    for {
        msg := <-c
        fmt.Println("msg : ", msg)
    }
}*/

func PrintFS( c <-chan string  ){
        for{
            msg:=<-c
            fmt.Println(msg)
        }
 }

func main() { 
    c := make(chan string) // c <- writing(c on left side) && -> c reading (c on right side) 
    go First(c)
    go Second(c)
    fmt.Println("test")
    go PrintFS(c)
    var input string
    fmt.Scanln(&input)
}