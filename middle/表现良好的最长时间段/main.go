package main

import "fmt"

func longestWPI(hours []int) int {
	n := len(hours)
	s := make([]int, n+1)
	stk := make([]int, 0)
	s[0] = 0
	stk = append(stk, 0)
	for i := 1; i <= n; i++ {
		if hours[i-1] > 8 {
			s[i] = s[i-1] + 1
		} else {
			s[i] = s[i-1] - 1
		}
		if len(stk) == 0 || s[stk[len(stk)-1]] > s[i] {
			stk = append(stk, i)
		}
	}
	// fmt.Printf("s: %v|stk: %v\n", s, stk)
	ans := 0
	for r := n; r > 0; r-- {
		for len(stk) != 0 && s[r] > s[stk[len(stk)-1]] {
			l := stk[len(stk)-1]
			stk = stk[0 : len(stk)-1]
			ans = max(ans, r-l)
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

func main() {
	hours := []int{9, 9, 6, 0, 6, 6, 9}
	res := longestWPI(hours)
	fmt.Printf("%v\n", res)
}
