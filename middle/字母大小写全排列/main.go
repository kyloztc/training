package main

import "unicode"

func letterCasePermutation(s string) (ans []string) {
	var dfs func([]byte, int)
	dfs = func(s []byte, pos int) {
		for pos < len(s) && unicode.IsDigit(rune(s[pos])) {
			pos++
		}
		if pos == len(s) {
			ans = append(ans, string(s))
			return
		}
		dfs(s, pos+1)
		s[pos] ^= 32
		dfs(s, pos+1)
	}
	dfs([]byte(s), 0)
	return
}
