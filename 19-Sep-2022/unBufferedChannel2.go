package main
 
import (
  "fmt"
  "time"
)

func main(){
	a:=1
	b:=2

	operationDone := make(chan bool)
	//unbuffered channel
	// make a channel that consumes and o/p bool
	go func() {
		fmt.Println(" process begin : ")
		b = a * b
		fmt.Println("wait for : ")
		time.Sleep(time.Second * 1)
		//operationDone <- true
		operationDone <- false

	}()
	time.Sleep(time.Second * 2)
	<- operationDone // channer waits till value is assigned to it, here in the anonymous func we assigned boolean 
	a = b * b
	fmt.Printf("a = %d, b=%d", a,b)
}