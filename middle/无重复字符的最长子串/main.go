package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	max := 0
	helper := make(map[string]int)
	for right < len(s) {
		rightChar := s[right : right+1]
		index, ok := helper[rightChar]
		if ok && (index+1) > left {
			left = index + 1
		}
		helper[rightChar] = right
		//fmt.Printf("right: %v|char: %v|existIndex: %v\n", right, rightChar, index)
		//fmt.Printf("left: %v|right: %v\n", left, right)
		currentLen := right - left + 1
		if currentLen > max {
			max = currentLen
		}
		right += 1
	}
	return max
}

func main() {
	str := "aabaab!bb"
	max := lengthOfLongestSubstring(str)
	fmt.Printf("max: %v\n", max)
}
