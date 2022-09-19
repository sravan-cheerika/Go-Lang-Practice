package main
import (
    "fmt"
    "time"
)
func main() {
    numbers := make(chan int, 3)
    go func() {

        fmt.Println("reading 1 :", <-numbers)
        fmt.Println("reading 2 :", <-numbers)
        /*// does not block
        fmt.Println("writing 1")
        numbers <- 1
        // does not block
        fmt.Println("writing 2")
        numbers <- 2
        // blocks after writing
        fmt.Println("writing 3")
        numbers <- 3
        // waiting to be unblocked
        fmt.Println("writing 4")
        numbers <- 4
        fmt.Println("writing 5")
        numbers <- 5 */
    }()
    // wait for all writes inside the go routine till it blocks
    time.Sleep(1 * time.Second)
    numbers<-1
//    fmt.Println("reading 1:", <-numbers) // does not block
    time.Sleep(1 * time.Second)
    numbers<-2
//    fmt.Println("reading 2:", <-numbers) // does not block
    /*fmt.Println("reading 1:", <-numbers) // does not block
    fmt.Println("reading 2:", <-numbers) // does not block
    fmt.Println("reading 3:", <-numbers) // does not block
    fmt.Println("reading 4:", <-numbers) // blocks (waits for channel write)
    fmt.Println("reading 5", <-numbers)*/

}