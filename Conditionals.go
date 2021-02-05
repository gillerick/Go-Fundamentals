package main

import (
	"fmt"
)

func main() {
	var x int
	x = 10
	if x < 20 {
		fmt.Printf("Yes, %v is less than 20", x)
	} else {
		fmt.Printf("No, %v is greater than 20", x)
	}

}
