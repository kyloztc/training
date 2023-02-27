package main

func movesToMakeZigzag(nums []int) int {
	return min(getCount(0, nums), getCount(1, nums))
}

func getCount(postion int, nums []int) int {
	res := 0
	for i := postion; i < len(nums); i += 2 {
		optTime := 0
		if i-1 >= 0 {
			optTime = max(nums[i]-nums[i-1]+1, optTime)
		}
		if i+1 < len(nums) {
			optTime = max(nums[i]-nums[i+1]+1, optTime)
		}
		res += optTime
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
