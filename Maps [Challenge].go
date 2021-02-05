package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Needles and pins Needles and pins Sew me a sail to catch me the wind"
	words := strings.Fields(text)
	//Creating an empty memo map to store words, as keys
	counts := map[string]int{}

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
}
