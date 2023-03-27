package main

func countSubstrings(s string, t string) int {
	n := len(s)
	m := len(t)
	dpl := make([][]int, n + 1)
	dpr := make([][]int, n + 1)
	for i := 0; i <= n; i ++ {
		dpl[i] = make([]int, m + 1)
		dpr[i] = make([]int, m + 1)
	}
	for i := 0; i < n; i ++ {
		for j := 0; j < m; j ++ {
			if s[i] == t[j] {
				dpl[i+1][j+1] = dpl[i][j] + 1
			} else {
				dpl[i+1][j+1] = 0
			}
		}
	}
	for i := n-1; i >= 0; i -- {
		for j := m - 1; j >=0; j -- {
			if s[i] == t[j] {
				dpr[i][j] = dpr[i+1][j+1] + 1
			} else {
				dpr[i][j] = 0
			}
		}
	}
	ans := 0
	for i := 0; i < n; i ++ {
		for j := 0; j < m; j ++ {
			if s[i] != t[j] {
				ans += (dpl[i][j] + 1) * (dpr[i+1][j+1] + 1)
			}
		}
	}
	return ans
}