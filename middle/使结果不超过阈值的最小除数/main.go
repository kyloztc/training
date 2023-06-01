package main

import "fmt"

func smallestDivisor(nums []int, threshold int) int {
	left := 1
	right := 0
	for _, num := range nums {
		right = max(right, num)
	}
	for left < right {
		mid := left + (right-left)/2
		// fmt.Printf("left: %v|right: %v\n", left, right)
		if getCount(nums, mid) > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func getCount(nums []int, divisor int) int {
	sum := 0
	for _, num := range nums {
		sum += (num + divisor - 1) / divisor
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 5, 9}
	threshold := 6
	res := smallestDivisor(nums, threshold)
	fmt.Printf("res: %v\n", res)
}
