// Generate the Series 1A2B3C4D5E....... using  goroutine and channel
// .Routine 1 shall generate A,B C D and routine 2 shall generate 1 ,2,3,4,5,6.......
// Both this output join together in the output

package main

import (
	"fmt"
//	"sync"
	"time"
)

// func alphabet() {
// 	for ch := 65; ch < 90; ch++ {
// 		fmt.Printf("%c ", ch)
// 	}
// 	//fmt.Println()
// }
// func numeric(ch chan int) {
// 	for i := 1; i <= 20; i++ {
// 		fmt.Printf("num", i)
// 	}
// }

/*var wg sync.WaitGroup

func numeric(ch chan int) {
	for i := 1; i <= 4; i++ {
		ch <- i
		time.Sleep(1 * time.Millisecond)
	}
	wg.Done()
}

func alphabet(ch chan rune) {
	for j := 'a'; j <= 'c'; j++ {
		ch <- j
		time.Sleep(1 * time.Millisecond)
	}
	wg.Done()
}*/

func main() {
	//wg.Add(2) // max num of functions to execute 

	ch,ch2 := make(chan rune), make(chan int)

	go func() {

        for j := 'A'; j <= 'Z'; j++ {
            ch <- j
			time.Sleep(1* time.Millisecond)
        }

    }()

	go func() {
        for i := 1; i < 35; i++ {
            ch2 <- i
			time.Sleep(1* time.Millisecond)
        }

    }()

	go func() {

        for {
            fmt.Print(string(<-ch))
			fmt.Print(<-ch2)
        }

    }() 
	time.Sleep(time.Second)
	//anynoums function call
	/*for i := 0; i < 6; i++ {
		select {
		case msg1 := <-ch:
			fmt.Print("", string(msg1))
		case msg2 := <-ch2:
			fmt.Print("", msg2)
		}
	}*/
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	//wg.Wait()
}
