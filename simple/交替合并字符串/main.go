package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	ptr1 := 0
	ptr2 := 0
	len1 := len(word1)
	len2 := len(word2)
	for ptr1 < len1 && ptr2 < len2 {
		res += string(word1[ptr1])
		res += string(word2[ptr2])
		ptr1 += 1
		ptr2 += 1
	}
	if ptr1 < len1 {
		res += word1[ptr1:]
	}
	if ptr2 < len2 {
		res += word2[ptr2:]
	}
	return res
}

func main() {
	word1 := "ab"
	word2 := "pqrs"
	res := mergeAlternately(word1, word2)
	fmt.Printf("res: %v\n", res)
}
