package main
import (
    "fmt"
    "sync"
    "time"
    "math/rand"
    )
var wg sync.WaitGroup
var counter int // Common Resource

func increment( s string) {
     rand.Seed(time.Now().UnixNano())
    for i := 0; i < 20; i++ {
        x := counter
        x++
        time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
        counter = x
        fmt.Println(s, i, "Counter:", counter)
    }
    wg.Done()
}
//go run -race main.go 
func main() {
    wg.Add(2)
    go increment("Anamika")
    go increment("Monika")
    wg.Wait()
}