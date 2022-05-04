package main

import "fmt"

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	generateHelper(n, 0, 0, "", &ans)
	return ans
}

func generateHelper(n int, left int, right int, str string, ans *[]string) {
	if left >= n {
		for i := 0; i < n-right; i++ {
			str += ")"
		}
		//fmt.Printf("str: %v\n", str)
		*ans = append(*ans, str)
		return
	}
	tmp := str
	tmp += "("
	generateHelper(n, left+1, right, tmp, ans)
	if left > right {
		tmp = str
		tmp += ")"
		generateHelper(n, left, right+1, tmp, ans)
	}
}

func main() {
	n := 3
	ans := generateParenthesis(n)
	fmt.Printf("ans: %v\n", ans)
}
