package main

import ( 
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello world")
    args:=os.Args
    argsV:=os.Args[1:]
    fmt.Println(args,  "  ", argsV)
}
