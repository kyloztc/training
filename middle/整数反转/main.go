package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}
	res := 0
	for x > 0 {
		lowestNum := x % 10
		res = res*10 + lowestNum
		x /= 10
		if res > math.MaxInt32 {
			return 0
		}
	}
	if isNegative {
		return -res
	}
	return res
}

func main() {
	num := -1230
	res := reverse(num)
	fmt.Printf("res: %v\n", res)
}
