package main

import "fmt"

func nextPermutation(nums []int) {
	numsLen := len(nums)
	smallIndex := numsLen - 2
	for smallIndex >= 0 && nums[smallIndex] >= nums[smallIndex+1] {
		smallIndex -= 1
	}
	if smallIndex >= 0 {
		largeIndex := numsLen - 1
		for largeIndex >= 0 && nums[largeIndex] <= nums[smallIndex] {
			largeIndex -= 1
		}
		nums[smallIndex], nums[largeIndex] = nums[largeIndex], nums[smallIndex]
	}
	i := smallIndex + 1
	j := numsLen - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i += 1
		j -= 1
	}
}

func main() {
	nums := []int{
		1, 2, 3,
	}
	nextPermutation(nums)
	fmt.Printf("res: %v\n", nums)
}
