package main

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0
	findIndex := 0
	numsLen := len(nums)
	for findIndex < numsLen {
		if nums[findIndex] != val {
			nums[index] = nums[findIndex]
			index += 1
		}
		findIndex += 1
	}
	// fmt.Printf("nums: %v\n", nums)
	return index
}

func main() {
	nums := []int{
		0, 1, 2, 2, 3, 0, 4, 2,
	}
	val := 2
	res := removeElement(nums, val)
	fmt.Printf("res: %v\n", res)
}
