package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2

}

//go passes an integer into a function by value, meaning, the original declaration remains constant
//However, Go passes a slice or map by reference
func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 3)
	fmt.Println(values)

}
