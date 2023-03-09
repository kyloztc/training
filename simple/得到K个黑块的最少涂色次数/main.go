package main

func minimumRecolors(blocks string, k int) int {
	n := len(blocks)
	pre := make([]int, n+1)
	for i := 0; i < n; i ++ {
		if blocks[i] == 'W' {
			pre[i+1] = pre[i] + 1
		} else {
			pre[i+1] = pre[i]
		}
	}
	ans := n
	for i := 0; i < n; i ++ {
		endIndex := i + k - 1
		if endIndex+1 > n {
			break
		}
		ans = min(ans, pre[endIndex+1] - pre[i])
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}