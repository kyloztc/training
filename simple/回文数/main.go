package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	numStr := strconv.Itoa(x)
	left := 0
	right := len(numStr) - 1
	for left < right {
		if numStr[left] != numStr[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	res := isPalindrome(12)
	fmt.Printf("res: %v\n", res)
}
