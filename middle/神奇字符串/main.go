package main

import "fmt"

func magicalString(n int) int {
	if n < 4 {
		return 1
	}
	nums := make([]int, n)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 2
	numPtr := 3
	countPtr := 2
	ans := 1
	isOne := 1
	for numPtr < n {
		for i := 0; i < nums[countPtr]; i++ {
			if isOne == 1 {
				nums[numPtr] = 1
				ans += 1
			} else {
				nums[numPtr] = 2
			}
			numPtr++
			if numPtr >= n {
				break
			}
		}
		isOne = isOne * -1
		countPtr++
	}
	fmt.Printf("%v\n", nums)
	return ans
}

func main() {
	res := magicalString(10)
	fmt.Printf("res: %v\n", res)
}
