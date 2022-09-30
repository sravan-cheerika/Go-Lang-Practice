package main

import (
	"fmt"
	"sort"
)

func ArrayOfVariadic(ls ...string) {
	fmt.Println("Array of string before sorting: ", ls)

	sort.Strings(ls)

	fmt.Println("After sorting : ", ls)

}
