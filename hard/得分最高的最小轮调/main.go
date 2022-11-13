package main

import "fmt"

func bestRotation(nums []int) int {
	n := len(nums)
	helper := make([][]int, n)
	for i, num := range nums {
		helper[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j >= num {
				helper[i][j] = 1
			}
		}
	}
	ansK := 0
	ans := 0
	for k := 0; k < n; k++ {
		sum := 0
		for i := 0; i < n; i++ {
			newIndex := i - k
			if newIndex < 0 {
				newIndex += n
			}
			// fmt.Printf("newIndex: %v\n", newIndex)
			sum += helper[i][newIndex]
		}
		// fmt.Printf("k: %v|sum: %v\n", k, sum)
		if sum > ans {
			ans = sum
			ansK = k
		}
	}
	return ansK
}

func main() {
	nums := []int{2, 3, 1, 4, 0}
	res := bestRotation(nums)
	fmt.Printf("res: %v\n", res)
}
