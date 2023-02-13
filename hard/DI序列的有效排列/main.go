package main

import "fmt"

func numPermsDISequence(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		dp[0][i] = 1
	}
	for i := 1; i < n+1; i++ {
		if s[i-1] == 'D' {
			for j := n - i; j >= 0; j -- {
				
			}
		}
		for num := 0; num <= n; num++ {
			sum := 0
			if s[i-1] == 'D' {
				for j := num + 1; j <= n; j++ {
					sum += dp[i-1][j]
				}
			} else {
				for j := num - 1; j >= 0; j-- {
					sum += dp[i-1][j]
				}
			}
			fmt.Printf("i: %v|num: %v|sum: %v\n", i, num, sum)
			dp[i][num] = sum
		}
	}
	res := 0
	for i := 0; i < n+1; i++ {
		res += dp[n][i]
	}
	return res
}

func main() {
	s := "DID"
	res := numPermsDISequence(s)
	fmt.Printf("res: %v\n", res)
}
