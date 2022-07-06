package main

func strStr(haystack string, needle string) int {
	next := genNext(needle)
	n := len(haystack)
	m := len(needle)
	i, j := 0, 0
	for i < n && j < m {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		return i - j
	} else {
		return -1
	}

}

func genNext(needle string) []int {
	next := make([]int, len(needle))
	next[0] = -1
	if len(needle) == 1 {
		return next
	}
	i := 0
	j := 1
	for j < len(needle)-1 {
		if i == -1 || needle[i] == needle[j] {
			i++
			j++
			next[j] = i
		} else {
			i = next[i]
		}
	}
	return next
}

func main() {

}
