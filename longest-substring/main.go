package main

import (
	"fmt"
)

func main() {
	s := "hakans"
	fmt.Println(lengthOfLongestSubstringV2(s))
}

func lengthOfLongestSubstringV2(s string) int {
	runes := []rune(s)

	store := make(map[rune]int)
	var left, result int
	for right, value := range s {
		store[value] += 1
		for store[value] > 1 {
			l := runes[left]
			store[l] -= 1
			left += 1
		}
		result = max(result, right-left+1)

		right += 1
	}

	return result
}
