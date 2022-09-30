package main

import (
	"fmt"
	"debug_Test/subDir"
)

func main() {
	
	subDir.Navigate(234)

	for i:=0; i<3; i++{
		fmt.Println("Debug ",i)
	}
	
}
