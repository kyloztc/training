package main

func nthUglyNumber(n int) int {
	if n <= 1 {
		return n
	}
	nums := []int{1}
	ptr2 := 0
	ptr3 := 0
	ptr5 := 0
	for n > 1 {
		_min := min(min(2*nums[ptr2], 3*nums[ptr3]), 5*nums[ptr5])
		if 2*nums[ptr2] == _min {
			ptr2++
		}
		if 3*nums[ptr3] == _min {
			ptr3++
		}
		if 5*nums[ptr5] == _min {
			ptr5++
		}
		nums = append(nums, _min)
		n -= 1
	}
	return nums[len(nums)-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
