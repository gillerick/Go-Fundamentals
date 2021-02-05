package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	var max int
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	fmt.Println(max)

}
