package subDir

import (
	"fmt"
)

func Navigate(i int) {
	for j := 0; j < 2; j++ {
		fmt.Println("inside navigate ", j)
		fmt.Println("testing", i)
	}

}
