package main
import (
    "fmt"
    "sync"
    "time"
    "math/rand"
    "sync/atomic"
    )
var wg sync.WaitGroup
var counter int64 // Common Resource
//var mutex sync.Mutex
var t=0
func increment( s string) {
     rand.Seed(time.Now().UnixNano())
    for i := 0; i < 20; i++ {
        time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
        atomic.AddInt64(&counter, 1)
        fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) // access without race 
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