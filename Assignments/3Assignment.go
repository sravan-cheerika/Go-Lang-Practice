package main

import (
	"fmt"
	//	"math/big"
)

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
// ans[i] is the number of 1's in the binary representation of i.

func countBitofOne(n int) int {
	count :=0
	for i:=0; i < len(n); i++{
		if(n & (1 << i)){
			fmt.Println("inside if: ")
			count++;
		}
	}
	return count
	//return (n & 1) + countBitofOne(n>>1)
}

func main() {
	i := 4
	fmt.Println("size is ",)
	// n := int64(i)
	//fmt.Sprintf("%b", big.NewInt(4))
	fmt.Println("count bits of 1 : ", countBitofOne(i))
}
