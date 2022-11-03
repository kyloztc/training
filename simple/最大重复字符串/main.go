package main

func maxRepeating(sequence string, word string) int {
	n, m := len(sequence), len(word)
	if n < m {
		return 0
	}
	ans := 0
	dp := make([]int, n)
	for i := m - 1; i < n; i++ {
		if sequence[i-m+1:i+1] == word {
			if i == m-1 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-m] + 1
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}
	return ans
}
