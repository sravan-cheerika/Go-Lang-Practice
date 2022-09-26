// Create Two goroutine .... Routine 1 should generate the Series from 1-100 and
// routine 2 check whether its numbers are prime number or not and get it print

package main

import (
	"fmt"
)

// Generate natural seri number: 2,3,4,...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter: delete the number which is divisible by a prime number to find prime number
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // returns channel i.e, a slice of Int values
	for i := 1; i <= 100; i++ {
		prime := <-ch

		fmt.Printf("%v: %v\n", i, prime)
		ch = PrimeFilter(ch, prime)

	}
}
