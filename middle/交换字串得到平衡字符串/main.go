package main

func balancedString(s string) int {
	count := make(map[byte]int)
	n := len(s) / 4
	for i := 0; i < len(s); i++ {
		count[s[i]] += 1
	}
	if check(count, n) {
		return 0
	}
	res := len(s)
	r := 0
	for l, c := range s {
		for r < len(s) && !check(count, n) {
			count[s[r]] -= 1
			r += 1
		}
		if !check(count, n) {
			break
		}
		res = min(res, r-l)
		count[byte(c)] += 1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func check(cnt map[byte]int, n int) bool {
	if cnt['Q'] > n || cnt['W'] > n || cnt['E'] > n || cnt['R'] > n {
		return false
	}
	return true
}
