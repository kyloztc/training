package main

func maximalNetworkRank(n int, roads [][]int) int {
	connect := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		connect[i] = make([]int, n)
	}
	for _, road := range roads {
		connect[road[0]][road[1]] = 1
		connect[road[1]][road[0]] = 1
		degree[road[0]] += 1
		degree[road[1]] += 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, degree[i]+degree[j]-connect[i][j])
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
