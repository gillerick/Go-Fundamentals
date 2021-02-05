package main

import "fmt"

func main() {

	loons := []string{"bugs", "daffy", "taz"}
	for i := 0; i < len(loons); i++ {
		fmt.Printf("%v %v\n", i+1, loons[i])
	}

	//Slice the loons from 1
	fmt.Println(loons[1:])

	//Array containing integers
	numbers := []int{2, 4, 5, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}
