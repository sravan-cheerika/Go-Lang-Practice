package main

import (
	"fmt"
	"sync"
)

func main() {
	bufferedChan := make(chan string, 2) //buffer of 3
	var wg sync.WaitGroup
	wg.Add(2) // capacity of parallel func execution
	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
		bufferedChan <- "fourth"
		fmt.Println("Sent 4th")
		bufferedChan <- "fifth"
		fmt.Println("Sent 5th")
		wg.Done()
	}()
	//time.Sleep(time.Duration(time.Second))
	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println("receive:", firstRead)
		secondRead := <-bufferedChan
		fmt.Println("receive:", secondRead)
		thirdRead := <-bufferedChan
		fmt.Println("receive:", thirdRead)
		fourthRead := <-bufferedChan
		fmt.Println("receive:", fourthRead)
		fifthRead := <-bufferedChan
		fmt.Println("receive:", fifthRead)
		wg.Done()
	}()
	wg.Wait()
}
