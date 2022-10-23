package main

import "fmt"

func maxAliveYear(birth []int, death []int) int {
	start := 1900
	end := 2000
	peopleNum := make([]int, end-start+2)
	for i := 0; i < len(birth); i++ {
		peopleNum[birth[i]-start] += 1
		peopleNum[death[i]-start+1] -= 1
	}
	_sum := 0
	_max := 0
	maxYear := 0
	for i := 0; i < len(peopleNum); i++ {
		_sum += peopleNum[i]
		if _sum > _max {
			_max = _sum
			maxYear = i
		}
	}
	return maxYear + start
}

func main() {
	birth := []int{1900, 1901, 1950}
	death := []int{1948, 1951, 2000}
	res := maxAliveYear(birth, death)
	fmt.Printf("res: %v\n", res)
}
