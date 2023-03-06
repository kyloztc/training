package main
func minimumDeletions(s string) int {
	n := len(s)
	a := 0
	b := 0
	for i := 0; i < n; i ++ {
		if s[i] == 'a' {
			a += 1
		} else {
			b = max(a, b) + 1
		}
	}
	return n - max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}