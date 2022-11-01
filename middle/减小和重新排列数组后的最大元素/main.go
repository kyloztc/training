package main

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	count := make([]int, n+1)
	for _, num := range arr {
		count[min(num, n)]++
	}
	miss := 0
	for _, num := range count[1:] {
		if num == 0 {
			miss++
		} else {
			miss -= min(num-1, miss)
		}
	}
	return n - miss
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
