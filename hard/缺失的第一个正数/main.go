package main

import "fmt"

func firstMissingPositive(nums []int) int {
	_max := len(nums) + 1
	for i, v := range nums {
		if v <= 0 {
			nums[i] = _max
		}
	}
	for _, v := range nums {
		_v := v
		if _v < 0 {
			_v = -_v
		}
		if _v < _max && nums[_v-1] > 0 {
			nums[_v-1] = -nums[_v-1]
		}
	}
	res := _max
	for i, v := range nums {
		fmt.Printf("v: %v\n", v)
		if v > 0 {
			res = i + 1
			break
		}
	}
	return res
}

func main() {
	nums := []int{
		1, 1,
	}
	res := firstMissingPositive(nums)
	fmt.Printf("res: %v\n", res)
}
