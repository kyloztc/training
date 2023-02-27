package main

import (
	"math"
	"sort"
)

func minTaps(n int, ranges []int) int {
	gap := make([][2]int, len(ranges))
	for i, r := range ranges {
		start := max(0, i-r)
		end := min(n, i + r)
		gap[i] = [2]int{start, end}
	}
	sort.Slice(gap, func(i, j int) bool {
		return gap[i][0] <= gap[j][0]
	})
	dp := make([]int, n + 1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for _, g := range gap {
		start := g[0]
		end := g[1]
		if dp[start] == math.MaxInt {
			return -1
		}
		for i := start; i <= end; i ++ {
			dp[i] = min(dp[start] + 1, dp[i])
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}