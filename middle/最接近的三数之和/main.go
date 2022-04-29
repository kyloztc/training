package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	numLen := len(nums)
	ans := math.MaxInt32
	for first := 0; first < numLen; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := numLen - 1
		for second := first + 1; second < numLen; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			if second >= third {
				break
			}
			sum := nums[first] + nums[second] + nums[third]
			//fmt.Printf("sum: %v\n", sum)
			if sum == target {
				return target
			}
			if math.Abs(float64(ans-target)) > math.Abs(float64(sum-target)) {
				ans = sum
			}
			if sum > target {
				third -= 1
				second -= 1
			}
		}
	}
	return ans
}

func main() {
	nums := []int{1, 1, -1, -1, 3}
	target := -1
	ans := threeSumClosest(nums, target)
	fmt.Printf("ans: %v\n", ans)
}
