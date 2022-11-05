package main

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i, str := range strs {
		zeros := strings.Count(str, "0")
		ones := len(str) - zeros
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zeros && k >= ones {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
