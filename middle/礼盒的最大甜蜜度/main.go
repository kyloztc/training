package main

import "sort"

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	left := 0
	right := price[len(price)-1] - price[0]
	for left < right {
		mid := (left + right + 1) / 2
		if check(price, k, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func check(price []int, k int, mid int) bool {
	pre := -1000000001
	cnt := 0
	for _, p := range price {
		if p-pre >= mid {
			cnt += 1
			pre = p
		}
	}
	return cnt >= k
}
