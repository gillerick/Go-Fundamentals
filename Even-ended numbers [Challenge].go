package main

import "fmt"

func main() {
	count := 0
	//For every pair of four-digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { //This ensures we don't count twice
			n := a * b

			//Check if a  * b even-ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}
	//Print count
	fmt.Println(count)
}
