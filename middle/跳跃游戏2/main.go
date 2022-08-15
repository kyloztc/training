package main

import "fmt"

func jump(nums []int) int {
	numLen := len(nums)
	res := make([]int, numLen)
	for i := 0; i < numLen; i++ {
		for j := i + 1; j <= i+nums[i] && j < numLen; j++ {
			_min := getMin(res[j], res[i]+1)
			if res[j] == 0 {
				res[j] = res[i] + 1
			} else {
				res[j] = _min
			}
		}
	}
	// for _, v := range res {
	// 	fmt.Printf("i: %v\n", v)
	// }
	return res[numLen-1]
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{
		2, 3, 0, 1, 4,
	}
	res := jump(nums)
	fmt.Printf("res: %v\n", res)
}
