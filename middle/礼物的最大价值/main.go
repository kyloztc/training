package main

func maxValue(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	dp := make([][]int, n + 1)
	for i := 0; i <= n; i ++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i ++ {
		for j := 1; j <= m; j ++ {
			dp[i][j] = max(dp[i-1][j] + grid[i-1][j-1], dp[i][j-1] + grid[i-1][j-1])
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}