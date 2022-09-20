package main

import (
	"fmt"
)

// Given two integers a and b, return the sum of the two integers
// without using the operators + and -.

func main() {
	a := 10
	b := 20

	fmt.Println("sum is :", add(a, b) )
}

func add(a int, b int) int {
	for i := 1; i <= b; i++ {
		a++
	}
	return a
}
