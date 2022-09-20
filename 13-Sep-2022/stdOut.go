package main

import "fmt"
import "os"


func main() {
    f,err:=os.Create("data.txt")
    f.Close()
//    fmt.Println("Hello world")
}
