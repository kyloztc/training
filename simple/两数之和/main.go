package main

import "fmt"

func twoSum(nums []int, target int) []int {
	helper := make(map[int][]int)
	for index, num := range nums {
		_, ok := helper[num]
		if !ok {
			helper[num] = make([]int, 0)
		}
		helper[num] = append(helper[num], index)
	}
	for index, num := range nums {
		rest := target - num
		indexList, ok := helper[rest]
		if !ok {
			continue
		}
		for _, i := range indexList {
			if i != index {
				return []int{index, i}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3,2,4}
	target := 6
	res := twoSum(nums, target)
	fmt.Printf("res: %v\n", res)
}
