package main

func numSmallerByFrequency(queries []string, words []string) []int {
	suffixSum := make([]int, 12)
	for _, s := range words {
		suffixSum[getCount(s)]++
	}
	for i := 10; i > 0; i-- {
		suffixSum[i-1] += suffixSum[i]
	}
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ans[i] = suffixSum[getCount(queries[i])+1]
	}
	return ans
}

func getCount(word string) int {
	pre := 'z'
	count := 0
	for _, c := range word {
		if c > pre {
			continue
		}
		if c < pre {
			count = 1
			pre = c
		} else {
			count += 1
		}
	}
	return count
}
