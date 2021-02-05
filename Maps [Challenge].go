package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Needles and pins Needles and pins Sew me a sail to catch me the wind"
	split_text := strings.Fields(text)
	//Creating an empty memo map to store words, as keys
	memo := map[string]int{}

	for i := 0; i < len(split_text); i++ {
		memo[split_text[i]] = 0
		//fmt.Printf("Text: %v Type (%T)", split_text[i], text[i])
	}

	for j := 0; j < len(split_text); j++ {
		_, ok := memo[split_text[j]]
		if ok {
			memo[split_text[j]] += 1
		}
		//fmt.Printf("Text: %v Type (%T)", split_text[i], text[i])
	}
	//fmt.Println(count)
	for k, v := range memo {
		fmt.Println(k, v)
	}
}
