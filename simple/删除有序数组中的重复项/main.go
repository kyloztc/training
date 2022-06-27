package main

import "fmt"

func removeDuplicates(nums []int) int {
	index := 1
	findIndex := 1
	count := 1
	for findIndex < len(nums) {
		for findIndex < len(nums) && nums[findIndex] == nums[findIndex-1] {
			findIndex += 1
		}
		if findIndex < len(nums) {
			nums[index] = nums[findIndex]
			index += 1
			findIndex += 1
			count += 1
		}
	}
	fmt.Printf("%v\n", nums)
	return count
}

func main() {
	nums := []int{
		0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
	}
	res := removeDuplicates(nums)
	fmt.Printf("res: %v\n", res)
}
