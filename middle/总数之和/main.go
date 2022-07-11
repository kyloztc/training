package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	current := make([]int, 0)
	combinationSumHelper(candidates, target, current, 0, &ans)
	return ans
}

func combinationSumHelper(candidates []int, target int, current []int, currentSum int, ans *[][]int) {
	if len(candidates) == 0 || currentSum > target {
		return
	}
	if currentSum == target {
		fmt.Printf("currnet: %v|current sum: %v\n", current, currentSum)
		_current := make([]int, 0)
		for _, value := range current {
			_current = append(_current, value)
		}
		*ans = append(*ans, _current)
		return
	}
	for index, value := range candidates {
		_currentSum := currentSum + value
		_current := current
		_current = append(_current, value)
		combinationSumHelper(candidates[index:], target, _current, _currentSum, ans)
	}
}

func main() {
	candidates := []int{
		2, 3, 5,
	}
	target := 8
	res := combinationSum(candidates, target)
	fmt.Printf("res: %v\n", res)
}
