package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	target := 0
	numLen := len(nums)
	res := make([][]int, 0)
	for first := 0; first < numLen; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		rest := target - nums[first]
		third := numLen - 1
		//fmt.Printf("rest: %v\n", rest)
		for second := first + 1; second < numLen; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > rest {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == rest {
				//fmt.Printf("first, second, third|%v|%v|%v", first, second, third)
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

func main() {
	nums := []int{
		-1, 0, 1, 2, -1, -4,
	}
	res := threeSum(nums)
	fmt.Printf("res: %v\n", res)
}
