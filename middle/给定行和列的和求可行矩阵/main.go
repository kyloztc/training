package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n := len(rowSum)
	m := len(colSum)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= res[i][j]
			colSum[j] -= res[i][j]
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
