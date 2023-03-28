package main

// https://leetcode.cn/problems/shortest-common-supersequence/solution/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-auy8z/
// 题解过于牛逼

func shortestCommonSupersequence(str1 string, str2 string) string {
	n := len(str1)
	m := len(str2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j]) + 1
			}
		}
	}
	ans := make([]byte, 0)
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if str1[i] == str2[j] {
			ans = append(ans, str1[i])
			i--
			j--
		} else if dp[i+1][j+1] == dp[i][j+1]+1 {
			ans = append(ans, str1[i])
			i--
		} else if dp[i+1][j+1] == dp[i+1][j]+1 {
			ans = append(ans, str2[j])
			j--
		}
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return str1[0:i+1] + str2[0:j+1] + string(ans)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
