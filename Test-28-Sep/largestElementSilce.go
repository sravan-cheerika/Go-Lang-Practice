package main

import "fmt"

func main(){
	array := []int {1,5,2,9,20,6,2}
	var largeNum, temp int

	for _, value := range array {
		if value > temp {
			temp = value
			largeNum = temp
		}
	}
	fmt.Println("largest value is :", largeNum)
}