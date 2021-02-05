package main

import "fmt"

func main() {
	n := 42
	fmt.Println(n)
	s := fmt.Sprintf("The string equivalent of n is %d", 42)
	fmt.Printf("s = %s (type %T)\n", s, s)

}
