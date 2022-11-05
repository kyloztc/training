package main

func minSwaps(nums []int) int {
	cnt := 0
	n := len(nums)
	pre := make([]int, n)
	ans := n
	for i, num := range nums {
		if num == 0 {
			if i == 0 {
				pre[i] = 1
			} else {
				pre[i] = pre[i-1] + 1
			}
		} else {
			if i != 0 {
				pre[i] = pre[i-1]
			}
			cnt++
		}
	}
	for i, num := range nums {
		if num == 1 {
			endIndex := (i + cnt - 1) % n
			count := 0
			if endIndex >= i {
				count = pre[endIndex] - pre[i]
			} else {
				count = pre[endIndex] + (pre[n-1] - pre[i])
			}
			ans = min(ans, count)
		}
	}
	if ans != n {
		return ans
	}
	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
