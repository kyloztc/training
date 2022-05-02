package main

import "fmt"

type stack struct {
	item      []string
	lastIndex int
}

func newStack() *stack {
	return &stack{
		item:      make([]string, 0),
		lastIndex: -1,
	}
}

func (s *stack) push(item string) {
	s.item = append(s.item, item)
	s.lastIndex += 1
}

func (s *stack) isEmpty() bool {
	if len(s.item) == 0 {
		return true
	}
	return false
}

func (s *stack) pop() string {
	if s.isEmpty() {
		return ""
	}
	res := s.item[s.lastIndex]
	s.item = s.item[0:s.lastIndex]
	s.lastIndex -= 1
	return res
}

func isValid(s string) bool {
	stack := newStack()
	pair := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for i := 0; i < len(s); i++ {
		current := s[i : i+1]
		switch current {
		case "(", "[", "{":
			stack.push(current)
		case ")", "]", "}":
			leftPart := stack.pop()
			checkPart := pair[current]
			if leftPart != checkPart {
				return false
			}
		}
	}
	return stack.isEmpty()
}

func main() {
	str := "}"
	res := isValid(str)
	fmt.Printf("res: %v\n", res)
}
