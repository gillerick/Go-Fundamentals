package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	//Begin with the assumption that the first item is the largest
	max := nums[0]
	//Use double-context to get the values and the indices
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)

}
