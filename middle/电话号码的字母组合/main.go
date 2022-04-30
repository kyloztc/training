package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return make([]string, 0)
	}
	letterMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	res := letterMap[digits[0:1]]
	for i := 1; i < len(digits); i++ {
		num := digits[i : i+1]
		helper := make([]string, 0)
		for _, str := range res {
			for _, digit := range letterMap[num] {
				helper = append(helper, str+digit)
			}
		}
		res = helper
	}
	return res
}

func main() {
	digits := "2"
	res := letterCombinations(digits)
	fmt.Printf("res: %v\n", res)
}
