package main

func shortestSubarray(nums []int, k int) int {
	numsLen := len(nums)
	preSum := make([]int, numsLen+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	ans := numsLen + 1
	helper := make([]int, 0)
	for i, curSum := range preSum {
		for len(helper) > 0 && curSum-preSum[helper[0]] >= k {
			ans = min(ans, i-helper[0])
			helper = helper[1:]
		}
		for len(helper) > 0 && preSum[helper[len(helper)-1]] >= curSum {
			helper = helper[:len(helper)-1]
		}
		helper = append(helper, i)
	}
	if ans < numsLen+1 {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
