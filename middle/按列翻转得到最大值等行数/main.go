package main

func maxEqualRowsAfterFlips(matrix [][]int) int {
	mapHelper := make(map[string]int)
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		str := make([]byte, m)
		for j := 0; j < m; j++ {
			if matrix[i][j]^matrix[i][0] == 0 {
				str[j] = '0'
			} else {
				str[j] = '1'
			}
		}
		mapHelper[string(str)]++
	}
	ans := 0
	for _, v := range mapHelper {
		if v > ans {
			ans = v
		}
	}
	return ans
}
