package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	wordLen := len(words[0])
	wordNum := len(words)
	windowLen := wordLen * wordNum
	wordMap := make(map[string]int)
	for _, word := range words {
		_, ok := wordMap[word]
		if !ok {
			wordMap[word] = 0
		}
		wordMap[word] = wordMap[word] + 1
	}
	for i := 0; i <= len(s)-windowLen; i++ {
		if containAll(s[i:i+windowLen], wordMap, wordLen, wordNum) {
			res = append(res, i)
		}
	}
	return res
}

func containAll(str string, wordMap map[string]int, wordLen int, wordNum int) bool {
	countMap := make(map[string]int)
	for i := 0; i < wordNum; i++ {
		subStr := str[i*wordLen : i*wordLen+wordLen]
		_, ok := wordMap[subStr]
		if !ok {
			countMap[subStr] = 0
		}
		countMap[subStr] = countMap[subStr] + 1
	}
	for key, value := range countMap {
		_v, ok := wordMap[key]
		if !ok || _v != value {
			return false
		}
	}
	return true
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{
		"word", "good", "best", "good",
	}
	res := findSubstring(s, words)
	fmt.Printf("res: %v\n", res)
}
