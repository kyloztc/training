package main

func countVowelStrings(n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}
	ans := 0
	for i := 0; i < 5; i++ {
		ans += dp[n-1][i]
	}
	return ans
}

func countVowelStrings1(n int) int {
	dp := [5]int{}
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < 5; j++ {
			dp[j] += dp[j-1]
		}
	}
	ret := 0
	for i := 0; i < 5; i++ {
		ret += dp[i]
	}
	return ret
}
