package main

import (
	"strconv"
)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	_max := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			currentNum, _ := strconv.Atoi(string(matrix[i][j]))
			if currentNum == 1 {
				if j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1] + 1
				}
				minWidth := dp[i][j]
				maxArea := minWidth
				_max = max(_max, maxArea)
				for k := i - 1; k >= 0; k-- {
					if dp[k][j] == 0 {
						break
					}
					minWidth = min(minWidth, dp[k][j])
					area := minWidth * (i - k + 1)
					_max = max(_max, area)
				}
			}
		}
	}
	return _max
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
