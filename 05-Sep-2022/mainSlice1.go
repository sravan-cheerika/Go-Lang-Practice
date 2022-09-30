package main

import "fmt"

func main() {
	names := []string{"sravan", "reddy", "cheerika", "hi"}
	for i, j := range names {
		fmt.Println(i, "  : ", j, len(names))
		fmt.Println()
	}
}
