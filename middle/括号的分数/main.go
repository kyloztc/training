package main

import "fmt"

func scoreOfParentheses(s string) int {
	stack := []int{0}
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" {
			stack = append(stack, 0)
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*v, 1)
		}
	}
	return stack[0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	s := "()"
	res := scoreOfParentheses(s)
	fmt.Printf("res: %v\n", res)
}
