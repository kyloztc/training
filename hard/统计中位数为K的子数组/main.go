package main

func countSubarrays(nums []int, k int) int {
	kIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			kIndex = i
			break
		}
	}
	count := make(map[int]int)
	count[0] = 1
	sum := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum += sign(nums[i], k)
		if i < kIndex {
			count[sum] += 1
		} else {
			pre0 := count[sum]
			pre1 := count[sum-1]
			ans += pre0 + pre1
		}
	}
	return ans
}

func sign(num int, k int) int {
	if num == k {
		return 0
	}
	if num < k {
		return -1
	}
	return 1
}
