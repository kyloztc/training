package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	nums := []int{
		1, 3, 5, 6,
	}
	target := 7
	res := searchInsert(nums, target)
	fmt.Printf("res: %v\n", res)
}
